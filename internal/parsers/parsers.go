package parsers

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/gocarina/gocsv"
	"github.com/xuri/excelize/v2"
)

func ParseCsvFile(filename string) (*[]internal.ScanRow, error) {
	start := time.Now()

	content, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scans := new([]internal.ScanRow)

	err = gocsv.Unmarshal(content, scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	log.Println(time.Since(start).Milliseconds())
	return scans, nil
}

func ParseExcelFile(filename string) (*[]internal.ScanRow, error) {
	start := time.Now()

	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found in excel document")
	}

	rows, err := f.Rows(sheets[0])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sb := strings.Builder{}

	for i := 0; rows.Next(); i++ {
		row, err := rows.Columns()
		if err != nil {
			return nil, fmt.Errorf("error parsing excel row %d: %s", i, err)
		}

		// TODO add warnings for this
		if len(row) < internal.ExpectedAmmountOfColumns {
			continue
		}

		sb.WriteString(fmt.Sprintf("%s\n", strings.Join(row, ",")))
	}

	scans := new([]internal.ScanRow)

	err = gocsv.UnmarshalString(sb.String(), scans)
	if err != nil {
		return nil, fmt.Errorf("gocsv: %s", err)
	}

	log.Println(time.Since(start).Milliseconds())
	return scans, nil
}
