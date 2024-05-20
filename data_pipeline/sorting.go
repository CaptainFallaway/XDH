package data_pipeline

import "sort"

// Sorts the grouping slice depeding on how many violations (times it has gone beyond accepted threshold)
// and if two grouping are the same they sort depending on index
func SortByViolations(groupings *[]Grouping, metal string) {
	sort.Slice(*groupings, func(left, right int) bool {
		leftViolation := (*groupings)[left].Violations[metal]
		rightViolation := (*groupings)[right].Violations[metal]
		leftIndex := (*groupings)[left].Index
		rightIndex := (*groupings)[right].Index

		if leftViolation == rightViolation {
			return leftIndex > rightIndex
		} else {
			return leftViolation > rightViolation
		}
	})
}
