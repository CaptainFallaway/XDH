package parsers

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"

	"github.com/CaptainFallaway/XDH/internal"
)

// readExcel reads an Excel file and converts it to a CSV string
func readExcel(filename string) (string, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return "", err
	}

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return "", fmt.Errorf("no sheets found in excel document")
	}

	rows, err := f.Rows(sheets[0])
	if err != nil {
		return "", err
	}

	defer func(rows *excelize.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("error closing rows: %s", err)
		}
	}(rows)

	sb := strings.Builder{}

	for i := 0; rows.Next(); i++ {
		row, err := rows.Columns()
		if err != nil {
			return "", fmt.Errorf("error parsing excel row %d: %s", i, err)
		}

		// TODO add warnings for this
		if len(row) < internal.ExpectedAmmountOfColumns {
			continue
		}

		sb.WriteString(fmt.Sprintf("%s\n", strings.Join(row, ",")))
	}

	return sb.String(), nil
}
