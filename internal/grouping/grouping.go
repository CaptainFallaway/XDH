package grouping

import (
	"github.com/CaptainFallaway/XDH/internal"
)

func MakeBoatGroupings(scans *[]internal.ScanRow) []internal.Grouping {
	boatMap := make(map[string][]internal.ScanRow)
	indexer := newIndexer() // Mitigate the randomness of maps

	// Pre-processing
	for _, scan := range *scans {
		indexer.AddIndexFor(scan.Boat)
		boatMap[scan.Boat] = append(boatMap[scan.Boat], scan)
	}

	grouping := make([]internal.Grouping, 0, len(boatMap))

	// Iterating over the map and creating the groupings for each boat with a builder
	for boatID, scans := range boatMap {
		builder := newGroupingBuilder(boatID)

		for _, scan := range scans {
			builder.Operators.Add(scan.Operator)
			builder.UnitSet.Add(scan.Units)
			builder.AppendScan(scan)
			builder.JustifyEarliestTime(scan.Time)
			builder.JustifyLatestTime(scan.Time)
			builder.CountViolations(scan)
		}

		grouping = append(grouping, builder.BuildGrouping(indexer.Indexes[boatID]))
	}

	return grouping
}
