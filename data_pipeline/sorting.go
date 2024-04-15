package data_pipeline

import "sort"

type Sortable interface {
	GetViolation(metal string) uint8
	GetIndex() uint32
}

// Sorts the grouping slice depeding on how many violations (times it has gone beyond accepted threshold)
// and if two grouping are the same they sort depending on index
func SortByViolations[T Sortable](groupings *[]T, metal string) {
	sort.Slice(*groupings, func(left, right int) bool {
		leftViolation := (*groupings)[left].GetViolation(metal)
		rightViolation := (*groupings)[right].GetViolation(metal)
		leftIndex := (*groupings)[left].GetIndex()
		rightIndex := (*groupings)[right].GetIndex()

		if leftViolation == rightViolation {
			return leftIndex > rightIndex
		} else {
			return leftViolation > rightViolation
		}
	})
}
