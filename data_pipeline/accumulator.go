package data_pipeline

import "fmt"

type accumulator struct {
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

func (a *accumulator) AppendScan(scan ScanRow) {
	if scan.Duration < ValidMinimumScanTime {
		a.InvalidScans = append(a.InvalidScans, scan)
	} else {
		a.Scans = append(a.Scans, scan)
	}
}

func (a *accumulator) AddErrorNote(err string) {
	a.ErrorNotes = append(a.ErrorNotes, err)
}

func (a *accumulator) JustifyEarilestTime(time string) {
	a.FirstDate = compareEarliestTimes(a.FirstDate, time)
}

func (a *accumulator) JustifyLatestTime(time string) {
	a.LastDate = compareLatestTimes(a.LastDate, time)
}

func (a *accumulator) CountViolations(scan ScanRow) {
	violationCount(scan, &a.ViolationCountMap)
}

func (a *accumulator) GetUnit() string {
	units := a.UnitSet.ToSlice()

	if len(units) > 1 {
		a.AddErrorNote(fmt.Sprintf("Multiple units detected: %v", len(units)))
	}

	return units[0]
}

func (a accumulator) BuildGrouping(index uint32) Grouping {
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

// We need to call this to instanciate the struct to avoid nil pointer dereference
func newAccumulator(boatID string) *accumulator {
	vcm := make(map[string]uint8, MetalPolicy.AmmountOfMetals)

	for _, metal := range MetalPolicy.Metals {
		vcm[metal] = 0
	}

	return &accumulator{
		UnitSet:           *newSet(),
		ViolationCountMap: vcm,
		BoatID:            boatID,
		Operators:         *newSet(),
	}
}
