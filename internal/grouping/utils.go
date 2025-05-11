package grouping

import (
	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/parsers"
)

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
