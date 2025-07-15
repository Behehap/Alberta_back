// internal/scheduler/scheduler.go
package scheduler

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Behehap/Alberta/internal/store"
)

// Task represents one unit of study that needs to be placed in the schedule.
type Task struct {
	BookID          int64
	IsPriority      bool
	RequiresPairing bool
	IsPaired        bool
}

// Slot represents a single 100-minute study block in the weekly grid.
type Slot struct {
	Day         time.Weekday
	BlockNumber int // 0=Morning, 1=Afternoon, 2=Evening
	IsAvailable bool
	BookID      int64 // 0 if empty
}

// Scheduler holds the dependencies needed for the scheduling algorithm.
type Scheduler struct {
	Store *store.Storage
}

// New creates a new Scheduler instance.
func New(s *store.Storage) *Scheduler {
	return &Scheduler{
		Store: s,
	}
}

// Generate creates a new weekly schedule for a student based on a template.
func (s *Scheduler) Generate(ctx context.Context, studentID int64, weeklyPlanID int64, templateID int64) error {
	// --- Step 1: Fetch all required data ---
	unavailableTimes, err := s.Store.UnavailableTimes.GetAllForStudent(ctx, studentID)
	if err != nil {
		return fmt.Errorf("could not get unavailable times: %w", err)
	}

	templateRules, err := s.Store.TemplateRules.GetAllForTemplate(ctx, templateID)
	if err != nil {
		return fmt.Errorf("could not get template rules: %w", err)
	}

	if len(templateRules) == 0 {
		return errors.New("no template rules found for the given template ID")
	}

	// --- Step 2: Generate the list of tasks to be scheduled ---
	var tasksToSchedule []Task
	for _, rule := range templateRules {
		for i := 0; i < rule.DefaultFrequency; i++ {
			task := Task{
				BookID:          rule.BookID,
				IsPriority:      strings.Contains(rule.SchedulingHints, "priority_first_block"),
				RequiresPairing: strings.Contains(rule.SchedulingHints, "contiguous_pair"),
			}
			tasksToSchedule = append(tasksToSchedule, task)
		}
	}

	// --- Step 3: Prepare the weekly grid ---
	scheduleGrid := make([][]Slot, 7)
	for i := 0; i < 7; i++ {
		scheduleGrid[i] = make([]Slot, 3)
		for j := 0; j < 3; j++ {
			scheduleGrid[i][j] = Slot{Day: time.Weekday(i), BlockNumber: j, IsAvailable: true}
		}
	}

	// Block out the unavailable times.
	for _, ut := range unavailableTimes {
		dayIndex := int(ut.DayOfWeek)
		// This logic assumes 3 blocks: Morning (0), Afternoon (1), Evening (2).
		// A more complex system could parse the times more granularly.
		startTime, _ := time.Parse("15:04:05", ut.StartTime)
		if startTime.Hour() < 12 { // Morning block
			scheduleGrid[dayIndex][0].IsAvailable = false
		} else if startTime.Hour() < 17 { // Afternoon block
			scheduleGrid[dayIndex][1].IsAvailable = false
		} else { // Evening block
			scheduleGrid[dayIndex][2].IsAvailable = false
		}
	}

	// --- Step 4: Run the placement engine ---
	tasksToSchedule = placePriorityTasks(tasksToSchedule, scheduleGrid)
	tasksToSchedule = placePairedTasks(tasksToSchedule, scheduleGrid)
	tasksToSchedule = placeGeneralTasks(tasksToSchedule, scheduleGrid)

	if len(tasksToSchedule) > 0 {
		return errors.New("not enough available slots to schedule all tasks")
	}

	// --- Step 5: Save the generated schedule to the database ---
	return s.saveSchedule(ctx, weeklyPlanID, scheduleGrid)
}

// placePriorityTasks places tasks marked as high-priority.
func placePriorityTasks(tasks []Task, grid [][]Slot) (remainingTasks []Task) {
	placedInDay := make(map[time.Weekday]bool)
	for _, task := range tasks {
		if !task.IsPriority {
			remainingTasks = append(remainingTasks, task)
			continue
		}
		placed := false
		for day := time.Saturday; day <= time.Friday; day++ {
			dayIndex := int(day) % 7
			if !placedInDay[day] && grid[dayIndex][0].IsAvailable {
				grid[dayIndex][0].BookID = task.BookID
				grid[dayIndex][0].IsAvailable = false
				placedInDay[day] = true
				placed = true
				break
			}
		}
		if !placed {
			remainingTasks = append(remainingTasks, task)
		}
	}
	return remainingTasks
}

// placePairedTasks places tasks that need to be scheduled back-to-back.
func placePairedTasks(tasks []Task, grid [][]Slot) (remainingTasks []Task) {
	var unpairedTasks []Task
	pairedTasks := make(map[int64][]Task)
	for _, task := range tasks {
		if task.RequiresPairing {
			pairedTasks[task.BookID] = append(pairedTasks[task.BookID], task)
		} else {
			unpairedTasks = append(unpairedTasks, task)
		}
	}

	for bookID, pair := range pairedTasks {
		if len(pair) < 2 {
			unpairedTasks = append(unpairedTasks, pair...)
			continue
		}
		placed := false
		for d := 0; d < 7; d++ {
			for b := 0; b < 2; b++ { // Only check blocks 0 and 1 for a following slot
				if grid[d][b].IsAvailable && grid[d][b+1].IsAvailable {
					grid[d][b].BookID = bookID
					grid[d][b].IsAvailable = false
					grid[d][b+1].BookID = bookID
					grid[d][b+1].IsAvailable = false
					placed = true
					break
				}
			}
			if placed {
				break
			}
		}
		if !placed {
			unpairedTasks = append(unpairedTasks, pair...)
		}
	}
	return unpairedTasks
}

// placeGeneralTasks fills the remaining slots with the rest of the tasks.
func placeGeneralTasks(tasks []Task, grid [][]Slot) (remainingTasks []Task) {
	rand.Shuffle(len(tasks), func(i, j int) { tasks[i], tasks[j] = tasks[j], tasks[i] })
	taskIndex := 0
	for d := 0; d < 7; d++ {
		for b := 0; b < 3; b++ {
			if grid[d][b].IsAvailable && taskIndex < len(tasks) {
				grid[d][b].BookID = tasks[taskIndex].BookID
				grid[d][b].IsAvailable = false
				taskIndex++
			}
		}
	}
	if taskIndex < len(tasks) {
		return tasks[taskIndex:]
	}
	return nil
}

// saveSchedule saves the generated grid to the database.
func (s *Scheduler) saveSchedule(ctx context.Context, weeklyPlanID int64, grid [][]Slot) error {
	tx, err := s.Store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback on error, commit will override this.

	dailyPlans := make(map[time.Weekday]int64)

	// Create daily plans and study sessions
	for d := 0; d < 7; d++ {
		dayHasSessions := false
		for b := 0; b < 3; b++ {
			if grid[d][b].BookID != 0 {
				dayHasSessions = true
				break
			}
		}

		if dayHasSessions {
			weeklyPlan, _ := s.Store.WeeklyPlans.Get(ctx, weeklyPlanID)
			planDate := weeklyPlan.StartDateOfWeek.AddDate(0, 0, d)
			dp := &store.DailyPlan{WeeklyPlanID: weeklyPlanID, PlanDate: planDate}
			err := s.Store.DailyPlans.Insert(ctx, dp) // This should be adapted to use the transaction
			if err != nil {
				return err
			}
			dailyPlans[time.Weekday(d)] = dp.ID
		}
	}

	for d := 0; d < 7; d++ {
		for b := 0; b < 3; b++ {
			if grid[d][b].BookID != 0 {
				dailyPlanID := dailyPlans[time.Weekday(d)]
				startTime, endTime := blockToTime(b)
				ss := &store.StudySession{
					DailyPlanID: dailyPlanID,
					LessonID:    1, // Placeholder: needs logic to select a lesson for the book
					StartTime:   startTime,
					EndTime:     endTime,
				}
				err := s.Store.StudySessions.Insert(ctx, ss) // This should also use the transaction
				if err != nil {
					return err
				}
			}
		}
	}

	return tx.Commit()
}

// blockToTime converts a block number to a start and end time string.
func blockToTime(blockNumber int) (string, string) {
	switch blockNumber {
	case 0: // Morning
		return "08:00:00", "09:40:00"
	case 1: // Afternoon
		return "13:00:00", "14:40:00"
	case 2: // Evening
		return "18:00:00", "19:40:00"
	default:
		return "", ""
	}
}
