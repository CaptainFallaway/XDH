package grouping

import (
	"fmt"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/models"
	"github.com/CaptainFallaway/XDH/internal/parsers"
)

type groupingBuilder struct {
	FirstDate         int64
	LastDate          int64
	Scans             []models.Scan
	ErrorNotes        []string
	UnitSet           set
	ViolationCountMap map[string]uint8
	BoatID            string
	Operators         set
}

func newGroupingBuilder(boatID string) *groupingBuilder {
	// Create the violation count map for the builder
	vcm := make(map[string]uint8, internal.MetalPolicy.AmmountOfMetals)

	// Initialize all of the metal values in the violation count map
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

// scanRowToScan converts the parser package dto to this scan type
func scanRowToScan(scan parsers.ScanRow) models.Scan {
	return models.Scan{
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
	scan.Valid = scan.Duration >= internal.ValidMinimumScanTime
	a.Scans = append(a.Scans, scan)
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

func (a *groupingBuilder) BuildGrouping(index int) models.Grouping {
	return models.Grouping{
		Index:      index,
		FirstDate:  a.FirstDate,
		LastDate:   a.LastDate,
		Unit:       a.getUnit(),
		ErrorNotes: a.ErrorNotes,
		Violations: a.ViolationCountMap,
		BoatID:     a.BoatID,
		Operators:  a.Operators.ToSlice(),
	}
}

func (a *groupingBuilder) GetScans() []models.Scan {
	// Count the amount of valid scans
	validScans := 0
	for _, scan := range a.Scans {
		if scan.Valid {
			validScans++
		}
	}

	if validScans < internal.MinimumAmmoutOfScans {
		a.AddErrorNote("Mindre än %d giltig scanner", internal.MinimumAmmoutOfScans)
	}

	return a.Scans
}
