package parsers

import (
	"fmt"
	"strings"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/gocarina/gocsv"
)

// Helper function for parsing the excel file
func marshalRow(row *[]string, header *[]string) (internal.ScanRow, error) {
	rowCsv := strings.Join(*header, ",") + "\n" + strings.Join(*row, ",")

	scan := make([]internal.ScanRow, 0, 1)

	err := gocsv.UnmarshalString(rowCsv, &scan)
	if err != nil {
		return internal.ScanRow{}, err
	}

	if len(scan) == 0 {
		return internal.ScanRow{}, fmt.Errorf("scan could not be loaded")
	}

	scanRow := scan[0]

	return scanRow, nil
}
