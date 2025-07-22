package scheduler

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/Behehap/Alberta/internal/store"
)

type Scheduler struct {
	Store *store.Storage
}

func NewScheduler(s *store.Storage) *Scheduler {
	return &Scheduler{
		Store: s,
	}
}

func generateTimeSlots(date time.Time, dayStartTime time.Time, dayEndTime time.Time, blockDuration time.Duration) []struct {
	Start time.Time
	End   time.Time
} {
	var slots []struct {
		Start time.Time
		End   time.Time
	}
	currentSlotStart := time.Date(date.Year(), date.Month(), date.Day(), dayStartTime.Hour(), dayStartTime.Minute(), dayStartTime.Second(), 0, time.Local)
	dayEndAdjusted := time.Date(date.Year(), date.Month(), date.Day(), dayEndTime.Hour(), dayEndTime.Minute(), dayEndTime.Second(), 0, time.Local)

	for currentSlotStart.Before(dayEndAdjusted) {
		slotEnd := currentSlotStart.Add(blockDuration)
		if slotEnd.After(dayEndAdjusted) {
			break
		}
		slots = append(slots, struct {
			Start time.Time
			End   time.Time
		}{Start: currentSlotStart, End: slotEnd})
		currentSlotStart = slotEnd
	}
	return slots
}

func (s *Scheduler) GenerateWeeklyPlan(
	ctx context.Context,
	studentID int64,
	weeklyPlanID int64,
	startDateOfWeek time.Time,
	totalStudyBlocksPerWeek int,
	unavailableTimes []*store.UnavailableTime,
	subjectFrequencies []*store.SubjectFrequency,
	templateRules []*store.TemplateRule,
) error {
	const blockDuration = 100 * time.Minute

	weeklyPlan, err := s.Store.WeeklyPlans.Get(ctx, weeklyPlanID)
	if err != nil {
		return fmt.Errorf("failed to retrieve weekly plan: %w", err)
	}

	dayStartDefault := time.Date(0, 1, 1, 8, 0, 0, 0, time.UTC)
	dayEndDefault := time.Date(0, 1, 1, 22, 0, 0, 0, time.UTC)

	if weeklyPlan.DayStartTime.Valid {
		dayStartDefault = weeklyPlan.DayStartTime.Time
	}

	availableSlotsPerDay := make(map[time.Weekday][]struct {
		Start time.Time
		End   time.Time
	})

	for i := 0; i < 7; i++ {
		currentDate := startDateOfWeek.AddDate(0, 0, i)
		currentWeekday := currentDate.Weekday()

		dailySlots := generateTimeSlots(currentDate, dayStartDefault, dayEndDefault, blockDuration)
		var filteredSlots []struct {
			Start time.Time
			End   time.Time
		}

		for _, slot := range dailySlots {
			isUnavailable := false
			for _, ut := range unavailableTimes {
				if ut.DayOfWeek == int(currentWeekday) || (!ut.IsRecurring && ut.DayOfWeek == -1) {
					var parsedUtStart, parsedUtEnd time.Time
					var err1, err2 error

					parsedUtStart, err1 = time.Parse(time.RFC3339, ut.StartTime)
					if err1 != nil {
						parsedUtStart, err1 = time.Parse("15:04:05", ut.StartTime)
					}

					parsedUtEnd, err2 = time.Parse(time.RFC3339, ut.EndTime)
					if err2 != nil {
						parsedUtEnd, err2 = time.Parse("15:04:05", ut.EndTime)
					}

					if err1 != nil || err2 != nil {
						continue
					}

					unavailableStart := time.Date(slot.Start.Year(), slot.Start.Month(), slot.Start.Day(), parsedUtStart.Hour(), parsedUtStart.Minute(), parsedUtStart.Second(), 0, time.Local)
					unavailableEnd := time.Date(slot.Start.Year(), slot.Start.Month(), slot.Start.Day(), parsedUtEnd.Hour(), parsedUtEnd.Minute(), parsedUtEnd.Second(), 0, time.Local)

					if slot.Start.Before(unavailableEnd) && slot.End.After(unavailableStart) {
						isUnavailable = true
						break
					}
				}
			}
			if !isUnavailable {
				filteredSlots = append(filteredSlots, slot)
			}
		}
		availableSlotsPerDay[currentWeekday] = filteredSlots
	}

	subjectsToSchedule := make(map[int64]int)
	for _, sf := range subjectFrequencies {
		subjectsToSchedule[sf.BookID] = sf.FrequencyPerWeek
	}

	rulesMap := make(map[int64]*store.TemplateRule)
	for _, rule := range templateRules {
		rulesMap[rule.BookID] = rule
	}

	scheduledCount := 0
	days := []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}

	for scheduledCount < totalStudyBlocksPerWeek {
		initialScheduledCount := scheduledCount
		for _, day := range days {
			currentDate := startDateOfWeek.AddDate(0, 0, int(day-startDateOfWeek.Weekday()+7)%7)

			dailyPlan, err := s.Store.DailyPlans.GetByWeeklyPlanAndDate(ctx, weeklyPlanID, currentDate)
			if err != nil {
				dailyPlan = &store.DailyPlan{
					WeeklyPlanID: weeklyPlanID,
					PlanDate:     currentDate,
				}
				err = s.Store.DailyPlans.Insert(ctx, dailyPlan)
				if err != nil {
					return fmt.Errorf("failed to create daily plan for %s: %w", currentDate.Format("2006-01-02"), err)
				}
			}

			slots := availableSlotsPerDay[day]
			if len(slots) == 0 {
				continue
			}

			var prioritizedBooks []int64
			var otherBooks []int64

			for bookID, remainingFreq := range subjectsToSchedule {
				if remainingFreq > 0 {
					rule, hasRule := rulesMap[bookID]
					if hasRule {
						if rule.PrioritySlot.Valid && rule.PrioritySlot.String == "first" {
							prioritizedBooks = append(prioritizedBooks, bookID)
							continue
						}
						if rule.TimePreference.Valid {
							if len(slots) > 0 {
								slotHour := slots[0].Start.Hour()
								if (rule.TimePreference.String == "morning" && slotHour < 12) ||
									(rule.TimePreference.String == "afternoon" && slotHour >= 12) {
									prioritizedBooks = append(prioritizedBooks, bookID)
									continue
								}
							}
						}
						if rule.ConsecutiveSessions.Valid && rule.ConsecutiveSessions.Bool {
							if len(slots) >= 2 && subjectsToSchedule[bookID] >= 2 {
								prioritizedBooks = append(prioritizedBooks, bookID)
								continue
							}
						}
					}
					otherBooks = append(otherBooks, bookID)
				}
			}

			sort.Slice(prioritizedBooks, func(i, j int) bool {
				return prioritizedBooks[i] < prioritizedBooks[j]
			})
			prioritizedBooks = append(prioritizedBooks, otherBooks...)

			for _, selectedBookID := range prioritizedBooks {
				if scheduledCount >= totalStudyBlocksPerWeek {
					break
				}
				if subjectsToSchedule[selectedBookID] <= 0 {
					continue
				}

				rule, hasRule := rulesMap[selectedBookID]

				if hasRule && rule.ConsecutiveSessions.Valid && rule.ConsecutiveSessions.Bool && subjectsToSchedule[selectedBookID] >= 2 && len(slots) >= 2 {

					studySession1 := &store.StudySession{
						DailyPlanID: dailyPlan.ID,
						BookID:      selectedBookID,
						StartTime:   slots[0].Start.Format("15:04:05"),
						EndTime:     slots[0].End.Format("15:04:05"),
						IsCompleted: false,
					}
					err = s.Store.StudySessions.Insert(ctx, studySession1)
					if err != nil {
						return fmt.Errorf("failed to insert consecutive study session 1: %w", err)
					}
					subjectsToSchedule[selectedBookID]--
					slots = slots[1:]
					availableSlotsPerDay[day] = slots
					scheduledCount++

					if scheduledCount < totalStudyBlocksPerWeek && subjectsToSchedule[selectedBookID] > 0 && len(slots) >= 1 {
						studySession2 := &store.StudySession{
							DailyPlanID: dailyPlan.ID,
							BookID:      selectedBookID,
							StartTime:   slots[0].Start.Format("15:04:05"),
							EndTime:     slots[0].End.Format("15:04:05"),
							IsCompleted: false,
						}
						err = s.Store.StudySessions.Insert(ctx, studySession2)
						if err != nil {
							return fmt.Errorf("failed to insert consecutive study session 2: %w", err)
						}
						subjectsToSchedule[selectedBookID]--
						slots = slots[1:]
						availableSlotsPerDay[day] = slots
						scheduledCount++
					}
					continue
				}

				if len(slots) > 0 && subjectsToSchedule[selectedBookID] > 0 {
					slot := slots[0]
					studySession := &store.StudySession{
						DailyPlanID: dailyPlan.ID,
						BookID:      selectedBookID,
						StartTime:   slot.Start.Format("15:04:05"),
						EndTime:     slot.End.Format("15:04:05"),
						IsCompleted: false,
					}
					err = s.Store.StudySessions.Insert(ctx, studySession)
					if err != nil {
						return fmt.Errorf("failed to insert study session: %w", err)
					}

					subjectsToSchedule[selectedBookID]--
					slots = slots[1:]
					availableSlotsPerDay[day] = slots
					scheduledCount++
				}
			}
		}

		if scheduledCount == initialScheduledCount && scheduledCount < totalStudyBlocksPerWeek {
			break
		}
	}

	return nil
}
