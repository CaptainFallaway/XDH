package parsers

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

func Parse(path string) ([]ScanRow, error) {
	if strings.HasSuffix(path, ".csv") {
		return ParseCsvFile(path)
	} else if strings.HasSuffix(path, ".xlsx") || strings.HasSuffix(path, ".xls") {
		return ParseExcelFile(path)
	} else {
		return nil, fmt.Errorf("parsers: file type not supported %s", path)
	}
}

func ParseCsvFile(path string) ([]ScanRow, error) {
	content, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os: %s", err)
	}

	scans := make([]ScanRow, 0)

	err = gocsv.Unmarshal(content, &scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	return scans, nil
}

func ParseExcelFile(path string) ([]ScanRow, error) {
	content, err := excelToCsv(path)
	if err != nil {
		return nil, fmt.Errorf("readExcel: %s", err)
	}

	scans := make([]ScanRow, 0)

	err = gocsv.UnmarshalString(content, &scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	return scans, nil
}
