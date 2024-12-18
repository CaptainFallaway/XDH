package grouping

import (
	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/parsers"
)

type set struct {
	vals map[string]struct{}
}

func newSet() *set {
	return &set{vals: make(map[string]struct{})}
}

func (s *set) Add(val string) {
	s.vals[val] = struct{}{}
}

func (s *set) ToSlice() []string {
	ret := make([]string, 0, len(s.vals))

	for k := range s.vals {
		ret = append(ret, k)
	}

	return ret
}

// This is a simple indexer that will index a string to a uint32.
// Basically like the Set object
type indexer struct {
	Indexes map[string]int
}

func newIndexer() *indexer {
	return &indexer{Indexes: make(map[string]int)}
}

func (i *indexer) AddIndexFor(val string) {
	if _, exists := i.Indexes[val]; !exists {
		i.Indexes[val] = len(i.Indexes)
	}
}

// Returns earliest time
func compareEarliestTimes(t1, t2 int64) int64 {
	// Since it'll be 0 because it's a zeroed date at first
	if t1 == 0 {
		return t2
	}

	if t1 < t2 {
		return t1
	} else {
		return t2
	}
}

// Returns latest time
func compareLatestTimes(t1, t2 int64) int64 {
	if t1 > t2 {
		return t1
	} else {
		return t2
	}
}

// VCM being the violation count map of the grouping
func violationCount(scan parsers.ScanRow, vcm *map[string]uint8) {
	if scan.Sn.Value >= internal.MetalPolicy.SnViolation {
		(*vcm)["Sn"] += 1
	}

	if scan.Cu.Value >= internal.MetalPolicy.CuViolation {
		(*vcm)["Cu"] += 1
	}

	if scan.Zn.Value >= internal.MetalPolicy.ZnViolation {
		(*vcm)["Zn"] += 1
	}

	if scan.Pb.Value >= internal.MetalPolicy.PbViolation {
		(*vcm)["Pb"] += 1
	}
}
