package parsers

import (
	"fmt"
	"github.com/CaptainFallaway/XDH/internal"
	"github.com/gocarina/gocsv"
	"os"
)

func ParseCsvFile(filename string) (*[]internal.ScanRow, error) {
	content, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os: %s", err)
	}

	scans := new([]internal.ScanRow)

	err = gocsv.Unmarshal(content, scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	return scans, nil
}

func ParseExcelFile(filename string) (*[]internal.ScanRow, error) {
	content, err := readExcel(filename)
	if err != nil {
		return nil, fmt.Errorf("readExcel: %s", err)
	}

	scans := new([]internal.ScanRow)

	err = gocsv.UnmarshalString(content, scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	return scans, nil
}
