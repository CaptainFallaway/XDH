package data_pipeline

import "fmt"

// The base accumulator using the identity as a generic

type baseAccumulator struct {
	FirstDate         string
	LastDate          string
	Scans             []ScanRow
	InvalidScans      []ScanRow
	ErrorNotes        []string
	UnitSet           set
	ViolationCountMap map[string]uint8
}

func (a *baseAccumulator) AppendScan(scan ScanRow) {
	if scan.Duration < ValidMinimumScanTime {
		a.InvalidScans = append(a.InvalidScans, scan)
	} else {
		a.Scans = append(a.Scans, scan)
	}
}

func (a *baseAccumulator) AddErrorNote(err string) {
	a.ErrorNotes = append(a.ErrorNotes, err)
}

func (a *baseAccumulator) JustifyEarilestTime(time string) {
	a.FirstDate = compareEarliestTimes(a.FirstDate, time)
}

func (a *baseAccumulator) JustifyLatestTime(time string) {
	a.LastDate = compareLatestTimes(a.LastDate, time)
}

func (a *baseAccumulator) CountViolations(scan ScanRow) {
	violationCount(scan, &a.ViolationCountMap)
}

func (a *baseAccumulator) GetUnit() string {
	units := a.UnitSet.ToSlice()

	if len(units) > 1 {
		a.AddErrorNote(fmt.Sprintf("Multiple units detected: %v", len(units)))
	}

	return units[0]
}

type boatIdAccumulator struct {
	baseAccumulator
	BoatID    string
	Operators set
}

func (a boatIdAccumulator) BuildGrouping(index uint32) BoatIDGrouping {
	return BoatIDGrouping{
		BaseGrouping: BaseGrouping{
			Index:        index,
			FirstDate:    a.FirstDate,
			LastDate:     a.LastDate,
			Unit:         a.GetUnit(),
			Scans:        a.Scans,
			InvalidScans: a.InvalidScans,
			ErrorNotes:   a.ErrorNotes,
			Violations:   a.ViolationCountMap,
		},
		BoatID:    a.BoatID,
		Operators: a.Operators.ToSlice(),
	}
}

type operatorAccumulator struct {
	baseAccumulator
	Operator string
	BoatIDs  set
}

func (a operatorAccumulator) BuildGrouping(index uint32) OperatorGrouping {
	return OperatorGrouping{
		BaseGrouping: BaseGrouping{
			Index:        index,
			FirstDate:    a.FirstDate,
			LastDate:     a.LastDate,
			Unit:         a.GetUnit(),
			Scans:        a.Scans,
			InvalidScans: a.InvalidScans,
			ErrorNotes:   a.ErrorNotes,
			Violations:   a.ViolationCountMap,
		},
		Operator: a.Operator,
		BoatIDs:  a.BoatIDs.ToSlice(),
	}
}

// We need to call this to instanciate the struct to avoid nil pointer dereference
func newBoatIdAccumulator(boatID string) *boatIdAccumulator {
	vcm := make(map[string]uint8, MetalPolicy.AmmountOfMetals)

	for _, metal := range MetalPolicy.Metals {
		vcm[metal] = 0
	}

	return &boatIdAccumulator{
		baseAccumulator: baseAccumulator{
			UnitSet:           *newSet(),
			ViolationCountMap: vcm,
		},
		BoatID:    boatID,
		Operators: *newSet(),
	}
}

// We need to call this to instanciate the struct to avoid nil pointer dereference
func newOperatorAccumulator(operator string) *operatorAccumulator {
	vcm := make(map[string]uint8, MetalPolicy.AmmountOfMetals)

	for _, metal := range MetalPolicy.Metals {
		vcm[metal] = 0
	}

	return &operatorAccumulator{
		baseAccumulator: baseAccumulator{
			UnitSet:           *newSet(),
			ViolationCountMap: vcm,
		},
		Operator: operator,
		BoatIDs:  *newSet(),
	}
}
