package grouping

import (
	"github.com/CaptainFallaway/XDH/internal"
)

func MakeBoatGroupings(scans *[]internal.ScanRow) []internal.Grouping {
	boatMap := make(map[string][]internal.ScanRow)
	indexer := newIndexer() // Mitigate the randomness of maps

	// Pre processing
	for _, scan := range *scans {
		indexer.AddIndexFor(scan.Boat)
		boatMap[scan.Boat] = append(boatMap[scan.Boat], scan)
		// log.Println(scan)
	}

	grouping := make([]internal.Grouping, 0, len(boatMap))

	// Iterating over the map and creating the groupings for each boat with a accumulator
	for boatID, scans := range boatMap {
		acc := newGroupingBuilder(boatID)

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
