package grouping

import (
	"fmt"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/parsers"
)

type groupingBuilder struct {
	FirstDate         int64
	LastDate          int64
	Scans             []internal.Scan
	InvalidScans      []internal.Scan
	ErrorNotes        []string
	UnitSet           set
	ViolationCountMap map[string]uint8
	BoatID            string
	Operators         set
}

func newGroupingBuilder(boatID string) *groupingBuilder {
	vcm := make(map[string]uint8, internal.MetalPolicy.AmmountOfMetals)

	for _, metal := range internal.MetalPolicy.Metals {
		vcm[metal] = 0
	}
	// Return a new grouping builder

	return &groupingBuilder{
		UnitSet:           *newSet(),
		ViolationCountMap: vcm,
		BoatID:            boatID,
		Operators:         *newSet(),
	}
}

// Converting parsers dto to domaing specific type
func scanRowToScan(scan parsers.ScanRow) internal.Scan {
	return internal.Scan{
		Reading:  scan.Reading,
		Duration: scan.Duration,
		Operator: scan.Operator,
		Date:     scan.Time.Unix,
		Pb:       scan.Pb.Value,
		Zn:       scan.Zn.Value,
		Cu:       scan.Cu.Value,
		Sn:       scan.Zn.Value,
		Violations: map[string]bool{
			"pb": scan.Pb.Value > internal.MetalPolicy.PbViolation,
			"zn": scan.Zn.Value > internal.MetalPolicy.ZnViolation,
			"cu": scan.Cu.Value > internal.MetalPolicy.CuViolation,
			"sn": scan.Sn.Value > internal.MetalPolicy.SnViolation,
		},
	}
}

func (a *groupingBuilder) AppendScan(scanRow parsers.ScanRow) {
	scan := scanRowToScan(scanRow)

	fmt.Println(scan)

	if scan.Duration < internal.ValidMinimumScanTime {
		a.InvalidScans = append(a.InvalidScans, scan)
	} else {
		a.Scans = append(a.Scans, scan)
	}
}

func (a *groupingBuilder) AddErrorNote(err string, args ...any) {
	a.ErrorNotes = append(a.ErrorNotes, fmt.Sprintf(err, args...))
}

func (a *groupingBuilder) JustifyEarliestTime(time parsers.Date) {
	a.FirstDate = compareEarliestTimes(a.FirstDate, time.Unix)
}

func (a *groupingBuilder) JustifyLatestTime(time parsers.Date) {
	a.LastDate = compareLatestTimes(a.LastDate, time.Unix)
}

func (a *groupingBuilder) CountViolations(scan parsers.ScanRow) {
	violationCount(scan, &a.ViolationCountMap)
}

// These functions below are specific since they will add warnings for the operator
// To keep note of when we are now building the grouping
func (a *groupingBuilder) getUnit() string {
	units := a.UnitSet.ToSlice()

	if len(units) > 1 {
		a.AddErrorNote("Multiple units detected: %v", len(units))
	}

	return units[0]
}

func (a *groupingBuilder) getValidScans() []internal.Scan {
	if len(a.Scans) < internal.MinimumAmmoutOfScans {
		a.AddErrorNote("Mindre än %d giltig scanner", internal.MinimumAmmoutOfScans)
	}

	return a.Scans
}

func (a *groupingBuilder) BuildGrouping(index int) internal.Grouping {
	return internal.Grouping{
		Index:        index,
		FirstDate:    a.FirstDate,
		LastDate:     a.LastDate,
		Unit:         a.getUnit(),
		Scans:        a.getValidScans(),
		InvalidScans: a.InvalidScans,
		ErrorNotes:   a.ErrorNotes,
		Violations:   a.ViolationCountMap,
		BoatID:       a.BoatID,
		Operators:    a.Operators.ToSlice(),
	}
}
