package grouping

import (
	"github.com/CaptainFallaway/XDH/internal/models"
	"github.com/CaptainFallaway/XDH/internal/parsers"
)

// MakeBoatGroupings takes in the parsed scan rows and builds the grouping objects.
//
// It also generates a Uid for each grouping.
func MakeBoatGroupings(scans []parsers.ScanRow) ([]models.Grouping, error) {
	boatMap := make(map[string][]parsers.ScanRow)
	indexer := newIndexer() // Mitigate the randomness of maps

	// Pre-processing
	for _, scan := range scans {
		indexer.AddIndexFor(scan.Boat)
		boatMap[scan.Boat] = append(boatMap[scan.Boat], scan)
	}

	groupings := make([]models.Grouping, 0, len(boatMap))

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

		grouping, err := builder.BuildGrouping(indexer.Indexes[boatID])
		if err != nil {
			return nil, err
		}

		groupings = append(groupings, grouping)
	}

	return groupings, nil
}
