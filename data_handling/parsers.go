package data_handling

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/xuri/excelize/v2"
)

func ParseCsv(filename string) ([]Scan, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var scans []Scan
	err = gocsv.UnmarshalBytes(bytes, &scans)
	if err != nil {
		return nil, err
	}

	return scans, nil
}

func ParseExcel(filename string) ([]Scan, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found in file")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, err
	}

	header := rows[0]
	var scans []Scan

	for i, row := range rows[1:] {
		scan, err := FromStringSlice(&row, &header)
		if err != nil {
			return nil, fmt.Errorf("error parsing row %d: %s", i, err)
		}

		scans = append(scans, scan)
	}

	return scans, nil
}

func FromStringSlice(row *[]string, header *[]string) (Scan, error) {
	row_csv := strings.Join(*header, ",") + "\n" + strings.Join(*row, ",")

	var scan []Scan
	err := gocsv.UnmarshalString(row_csv, &scan)
	if err != nil {
		return Scan{}, err
	}

	scan_item := scan[0]
	parser, err := time.Parse("1/2/06 15:04", scan_item.Time)
	if err != nil {
		return Scan{}, err
	}

	scan_item.Time = parser.Format("2006-01-02 15:04")

	return scan_item, nil
}
