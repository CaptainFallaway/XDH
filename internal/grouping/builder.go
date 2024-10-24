package grouping

import (
	"fmt"

	"github.com/CaptainFallaway/XDH/internal"
)

type groupingBuilder struct {
	FirstDate         internal.Date
	LastDate          internal.Date
	Scans             []internal.ScanRow
	InvalidScans      []internal.ScanRow
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

	return &groupingBuilder{
		UnitSet:           *newSet(),
		ViolationCountMap: vcm,
		BoatID:            boatID,
		Operators:         *newSet(),
	}
}

func (a *groupingBuilder) AppendScan(scan internal.ScanRow) {
	if scan.Duration < internal.ValidMinimumScanTime {
		a.InvalidScans = append(a.InvalidScans, scan)
	} else {
		a.Scans = append(a.Scans, scan)
	}
}

func (a *groupingBuilder) AddErrorNote(err string) {
	a.ErrorNotes = append(a.ErrorNotes, err)
}

func (a *groupingBuilder) JustifyEarilestTime(time internal.Date) {
	a.FirstDate = compareEarliestTimes(a.FirstDate, time)
}

func (a *groupingBuilder) JustifyLatestTime(time internal.Date) {
	a.LastDate = compareLatestTimes(a.LastDate, time)
}

func (a *groupingBuilder) CountViolations(scan internal.ScanRow) {
	violationCount(scan, &a.ViolationCountMap)
}

func (a *groupingBuilder) GetUnit() string {
	units := a.UnitSet.ToSlice()

	if len(units) > 1 {
		a.AddErrorNote(fmt.Sprintf("Multiple units detected: %v", len(units)))
	}

	return units[0]
}

func (a groupingBuilder) BuildGrouping(index int) internal.Grouping {
	return internal.Grouping{
		Index:        index,
		FirstDate:    a.FirstDate,
		LastDate:     a.LastDate,
		Unit:         a.GetUnit(),
		Scans:        a.Scans,
		InvalidScans: a.InvalidScans,
		ErrorNotes:   a.ErrorNotes,
		Violations:   a.ViolationCountMap,
		BoatID:       a.BoatID,
		Operators:    a.Operators.ToSlice(),
	}
}
