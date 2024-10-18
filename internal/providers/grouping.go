package data_pipeline

// A grouping is a colletion of info and scans for a specific boat
type Grouping struct {
	Index        uint32           `json:"index"`
	FirstDate    string           `json:"firstDate"`
	LastDate     string           `json:"lastDate"`
	Unit         string           `json:"unit"`
	Scans        []ScanRow        `json:"scans"`
	InvalidScans []ScanRow        `json:"invalidScans"`
	ErrorNotes   []string         `json:"errorNotes"`
	Violations   map[string]uint8 `json:"violations"`
	BoatID       string           `json:"boatID"`
	Operators    []string         `json:"operators"`
}

func MakeBoatGroupings(scans *[]ScanRow) []Grouping {
	boatMap := make(map[string][]ScanRow)
	indexer := newIndexer() // Mitigate the randomness of maps

	// Pre processing
	for _, scan := range *scans {
		indexer.AddIndexFor(scan.Boat)
		boatMap[scan.Boat] = append(boatMap[scan.Boat], scan)
	}

	grouping := make([]Grouping, 0, len(boatMap))

	// Iterating over the map and creating the groupings for each boat with a accumulator
	for boatID, scans := range boatMap {
		acc := newBoatModelBuilder(boatID)

		if len(scans) < 8 {
			acc.AddErrorNote("Less than 8 scans")
		}

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
