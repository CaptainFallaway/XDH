package data_pipeline

// This is also supposed to model what is going to show on the frontend
type BaseGrouping struct {
	Index        uint32
	FirstDate    string
	LastDate     string
	Unit         string
	Scans        []ScanRow
	InvalidScans []ScanRow        // Scans that are under 20 seconds which is the chalmars standard
	ErrorNotes   []string         // Notes that may be presented in the UI, get collected during grouping
	Violations   map[string]uint8 // A map that contains each metal and the amount of times it has been over the "safe" value
}

func (bg BaseGrouping) GetViolation(metal string) uint8 {
	return bg.Violations[metal]
}

func (bg BaseGrouping) GetIndex() uint32 {
	return bg.Index
}

type BoatIDGrouping struct {
	BaseGrouping
	BoatID    string
	Operators []string
}

type OperatorGrouping struct {
	BaseGrouping
	Operator string
	BoatIDs  []string
}

func GroupByBoat(scans *[]ScanRow) []BoatIDGrouping {
	boatMap := make(map[string][]ScanRow)
	indexer := newIndexer() // I do this to mitigate the randomness of maps

	// Pre processing
	for _, scan := range *scans {
		indexer.AddIndexFor(scan.Boat)
		boatMap[scan.Boat] = append(boatMap[scan.Boat], scan)
	}

	grouping := make([]BoatIDGrouping, 0, len(boatMap))

	// Iterating over the map and creating the groupings for each boat with a accumulator
	for boatID, scans := range boatMap {
		acc := newBoatIdAccumulator(boatID)

		for _, scan := range scans {
			acc.Operators.Add(scan.Operator)
			acc.UnitSet.Add(scan.Units)
			acc.AppendScan(scan)
			acc.JustifyEarilestTime(scan.Time)
			acc.JustifyLatestTime(scan.Time)
			acc.CountViolations(scan)
		}

		grouping = append(grouping, acc.BuildGrouping(indexer.Indexes[boatID]))
	}

	return grouping
}

func GroupByOperator(scans *[]ScanRow) []OperatorGrouping {
	operatorMap := make(map[string][]ScanRow)
	indexer := newIndexer() // I do this to mitigate the randomness of maps

	// Pre processing
	for _, scan := range *scans {
		indexer.AddIndexFor(scan.Operator)
		operatorMap[scan.Operator] = append(operatorMap[scan.Operator], scan)
	}

	grouping := make([]OperatorGrouping, 0, len(operatorMap))

	// Iterating over the map and creating the groupings for each boat with a accumulator
	for operator, scans := range operatorMap {
		acc := newOperatorAccumulator(operator)

		for _, scan := range scans {
			acc.BoatIDs.Add(scan.Boat)
			acc.UnitSet.Add(scan.Units)
			acc.AppendScan(scan)
			acc.JustifyEarilestTime(scan.Time)
			acc.JustifyLatestTime(scan.Time)
			acc.CountViolations(scan)
		}

		grouping = append(grouping, acc.BuildGrouping(indexer.Indexes[operator]))
	}

	return grouping
}
