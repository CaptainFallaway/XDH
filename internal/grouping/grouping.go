package grouping

import (
	"github.com/CaptainFallaway/XDH/internal/models"
	"github.com/CaptainFallaway/XDH/internal/parsers"
)

// MakeBoatGroupings takes in the parsed scan rows and builds the grouping objects.
//
// This is intentionally a static relation since the builder of this package
// Is bound to what is then stored in the database and operated on from the user.
func BomboclatConverter(scans []parsers.ScanRow) ([]models.Grouping, []models.Scan) {
	boatMap := make(map[string][]parsers.ScanRow)
	indexer := newIndexer() // Mitigate the randomness of maps

	// Pre-processing
	for _, scan := range scans {
		indexer.AddIndexFor(scan.Boat)
		boatMap[scan.Boat] = append(boatMap[scan.Boat], scan)
	}

	grouping := make([]models.Grouping, 0, len(boatMap))

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

	return grouping, builder.Scans
}
