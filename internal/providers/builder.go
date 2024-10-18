package data_pipeline

import (
	"fmt"

	"github.com/CaptainFallaway/XDH/internal"
)

type boatModelBuilder struct {
	FirstDate         string
	LastDate          string
	Scans             []ScanRow
	InvalidScans      []ScanRow
	ErrorNotes        []string
	UnitSet           set
	ViolationCountMap map[string]uint8
	BoatID            string
	Operators         set
}

func newBoatModelBuilder(boatID string) *boatModelBuilder {
	vcm := make(map[string]uint8, internal.MetalPolicy.AmmountOfMetals)

	for _, metal := range internal.MetalPolicy.Metals {
		vcm[metal] = 0
	}

	return &boatModelBuilder{
		UnitSet:           *newSet(),
		ViolationCountMap: vcm,
		BoatID:            boatID,
		Operators:         *newSet(),
	}
}

func (a *boatModelBuilder) AppendScan(scan ScanRow) {
	if scan.Duration < internal.ValidMinimumScanTime {
		a.InvalidScans = append(a.InvalidScans, scan)
	} else {
		a.Scans = append(a.Scans, scan)
	}
}

func (a *boatModelBuilder) AddErrorNote(err string) {
	a.ErrorNotes = append(a.ErrorNotes, err)
}

func (a *boatModelBuilder) JustifyEarilestTime(time string) {
	a.FirstDate = compareEarliestTimes(a.FirstDate, time)
}

func (a *boatModelBuilder) JustifyLatestTime(time string) {
	a.LastDate = compareLatestTimes(a.LastDate, time)
}

func (a *boatModelBuilder) CountViolations(scan ScanRow) {
	violationCount(scan, &a.ViolationCountMap)
}

func (a *boatModelBuilder) GetUnit() string {
	units := a.UnitSet.ToSlice()

	if len(units) > 1 {
		a.AddErrorNote(fmt.Sprintf("Multiple units detected: %v", len(units)))
	}

	return units[0]
}

func (a boatModelBuilder) BuildGrouping(index uint32) Grouping {
	return Grouping{
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
