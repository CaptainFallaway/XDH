package internal

import (
	"sort"

	"github.com/CaptainFallaway/XDH/internal/models"
)

// SortByViolations Sorts the grouping slice depending on how many violations (times it has gone beyond accepted threshold)
// and if two grouping are the same they sort depending on index
func SortByViolations(groupings []models.Grouping, metal string) {
	sort.Slice(groupings, func(left, right int) bool {
		leftViolation := groupings[left].Violations[metal]
		rightViolation := groupings[right].Violations[metal]
		leftIndex := groupings[left].Index
		rightIndex := groupings[right].Index

		if leftViolation == rightViolation {
			return leftIndex > rightIndex
		} else {
			return leftViolation > rightViolation
		}
	})
}
