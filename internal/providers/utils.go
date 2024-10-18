package data_pipeline

import (
	"time"

	"github.com/CaptainFallaway/XDH/internal"
)

type set struct {
	vals map[string]bool
}

func (s *set) Add(val string) {
	(*s).vals[val] = false
}

func (s *set) ToSlice() []string {
	ret := make([]string, 0, len(s.vals))

	for k := range s.vals {
		ret = append(ret, k)
	}

	return ret
}

func newSet() *set {
	return &set{vals: make(map[string]bool)}
}

// This is a simple indexer that will index a string to a uint32.
// Basically like the Set object
type indexer struct {
	Indexes map[string]uint32
}

// Add a index for a value.
// Basically just a counter set
func (i *indexer) AddIndexFor(val string) {
	if _, exists := i.Indexes[val]; !exists {
		i.Indexes[val] = uint32(len(i.Indexes))
	}
}

func newIndexer() *indexer {
	return &indexer{Indexes: make(map[string]uint32)}
}

// Returns earliest time
func compareEarliestTimes(t1, t2 string) string {
	pt1, err := time.Parse(internal.TimeFormat, t1)
	if err != nil {
		return t2
	}

	pt2, err := time.Parse(internal.TimeFormat, t2)
	if err != nil {
		return t1
	}

	if pt1.Unix() < pt2.Unix() {
		return pt1.Format(internal.TimeFormat)
	} else {
		return pt2.Format(internal.TimeFormat)
	}
}

// Returns latest time
func compareLatestTimes(t1, t2 string) string {
	pt1, err := time.Parse(internal.TimeFormat, t1)
	if err != nil {
		return t2
	}

	pt2, err := time.Parse(internal.TimeFormat, t2)
	if err != nil {
		return t1
	}

	if pt1.Unix() > pt2.Unix() {
		return pt1.Format(internal.TimeFormat)
	} else {
		return pt2.Format(internal.TimeFormat)
	}
}

func violationCount(scan ScanRow, vcm *map[string]uint8) {
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
