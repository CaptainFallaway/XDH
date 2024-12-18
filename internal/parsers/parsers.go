package parsers

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func ParseCsvFile(filename string) ([]ScanRow, error) {
	content, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os: %s", err)
	}

	scans := make([]ScanRow, 0)

	err = gocsv.Unmarshal(content, scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	return scans, nil
}

func ParseExcelFile(filename string) ([]ScanRow, error) {
	content, err := excelToCsv(filename)
	if err != nil {
		return nil, fmt.Errorf("readExcel: %s", err)
	}

	scans := make([]ScanRow, 0)

	err = gocsv.UnmarshalString(content, scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	return scans, nil
}
