package scheduler

import (
	"context"
	"fmt"
	"math"

	"github.com/Behehap/Alberta/internal/store"
)

type TemplateMatcher struct {
	Store *store.Storage
}

func NewTemplateMatcher(store *store.Storage) *TemplateMatcher {
	return &TemplateMatcher{
		Store: store,
	}
}

// FindClosestTemplate finds the template with total study blocks closest to target
func (tm *TemplateMatcher) FindClosestTemplate(ctx context.Context, gradeID, majorID int64, targetBlocks int) (*store.ScheduleTemplate, error) {
	templates, err := tm.Store.ScheduleTemplates.GetAll(ctx, gradeID, majorID)
	if err != nil {
		return nil, err
	}

	if len(templates) == 0 {
		return nil, nil
	}

	var closestTemplate *store.ScheduleTemplate
	minDiff := math.MaxInt32

	for _, template := range templates {
		diff := int(math.Abs(float64(template.TotalStudyBlocksPerWeek - targetBlocks)))
		if diff < minDiff {
			minDiff = diff
			closestTemplate = template
		}
	}

	return closestTemplate, nil
}

// CalculateAdjustedFrequencies adjusts frequencies based on template and target blocks
func (tm *TemplateMatcher) CalculateAdjustedFrequencies(
	ctx context.Context,
	template *store.ScheduleTemplate,
	selectedBookIDs []int64,
	targetBlocks int,
) (map[int64]int, error) {

	// For now, use simple equal distribution
	return tm.distributeFrequenciesEqually(selectedBookIDs, targetBlocks, ctx, template.ID)
}

// Simple equal distribution - we'll enhance this later
func (tm *TemplateMatcher) distributeFrequenciesEqually(bookIDs []int64, targetBlocks int, ctx context.Context, templateID int64) (map[int64]int, error) {
	numSubjects := len(bookIDs)

	if numSubjects > targetBlocks {
		return nil, fmt.Errorf("too many subjects (%d) for available blocks (%d)", numSubjects, targetBlocks)
	}

	freqMap := make(map[int64]int)

	// Calculate base frequency (integer division)
	baseFrequency := targetBlocks / numSubjects
	remainder := targetBlocks % numSubjects

	// Assign base frequency to all subjects
	for _, bookID := range bookIDs {
		freqMap[bookID] = baseFrequency
	}

	// Distribute remainder to first few subjects
	for i := 0; i < remainder; i++ {
		if i < len(bookIDs) {
			freqMap[bookIDs[i]]++
		}
	}

	return freqMap, nil
}

// Filter rules by selected books (helper function)
func filterRulesByBooks(rules []*store.TemplateRule, selectedBookIDs []int64) []*store.TemplateRule {
	selectedMap := make(map[int64]bool)
	for _, id := range selectedBookIDs {
		selectedMap[id] = true
	}

	var filtered []*store.TemplateRule
	for _, rule := range rules {
		if selectedMap[rule.BookID] {
			filtered = append(filtered, rule)
		}
	}
	return filtered
}

// Calculate total blocks from template rules
func calculateTotalBlocks(rules []*store.TemplateRule) int {
	total := 0
	for _, rule := range rules {
		total += rule.DefaultFrequency
	}
	return total
}

// Convert rules to frequency map
func convertRulesToFrequencyMap(rules []*store.TemplateRule) map[int64]int {
	freqMap := make(map[int64]int)
	for _, rule := range rules {
		freqMap[rule.BookID] = rule.DefaultFrequency
	}
	return freqMap
}
