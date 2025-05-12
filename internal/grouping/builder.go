package grouping

import (
	"fmt"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/models"
	"github.com/CaptainFallaway/XDH/internal/parsers"
	"github.com/google/uuid"
)

type groupingBuilder struct {
	FirstDate         int64
	LastDate          int64
	Scans             []models.Scan
	InvalidScans      []models.Scan
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

// getValidScans returns the valid scans and adds an error note if there are too few
func (a *groupingBuilder) getValidScans() []models.Scan {
	if len(a.Scans) < internal.MinimumAmmoutOfScans {
		a.AddErrorNote("Mindre än %d giltig scanner", internal.MinimumAmmoutOfScans)
	}

	return a.Scans
}

// BuildGrouping builds the grouping object from the builder.
// It also generates a Uid for each grouping.
func (a *groupingBuilder) BuildGrouping(index int) (models.Grouping, error) {
	uid, err := uuid.NewV7()

	return models.Grouping{
		Uid:          uid.String(),
		Index:        index,
		FirstDate:    a.FirstDate,
		LastDate:     a.LastDate,
		Unit:         a.getUnit(),
		ValidScans:   a.getValidScans(),
		InvalidScans: a.InvalidScans,
		ErrorNotes:   a.ErrorNotes,
		Violations:   a.ViolationCountMap,
		BoatID:       a.BoatID,
		Operators:    a.Operators.ToSlice(),
	}, err
}
