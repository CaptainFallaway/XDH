package data_pipeline

import (
	"fmt"
	"os"
	"strings"
	"time"

	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/xuri/excelize/v2"
)

type MetalValue struct {
	Value float64
	IsLod bool
}

func (n *MetalValue) UnmarshalCSV(val string) (err error) {
	if val, err := strconv.ParseFloat(val, 64); err == nil {
		n.Value = val
		n.IsLod = false
	} else {
		n.IsLod = true
	}
	return err
}

func (n *MetalValue) MarshalCSV() (string, error) {
	panic("Not implemented")
}

// Index,Reading No,Time,Type,Duration,Units,Sigma Value,Sequence,User1,Flags,Boat,Operator,User Login,Pb,Pb Error,Hg,Hg Error,W,W Error,Rb,Rb Error,Se,Se Error,Zn,Zn Error,Ag,Ag Error,Bk1,Bk1 Error,Mo,Mo Error,Sr,Sr Error,As,As Error,Cu,Cu Error,Ni,Ni Error,Fe,Fe Error,Cr,Cr Error,Bk2,Bk2 Error,Bk3,Bk3 Error,Bk4,Bk4 Error,Sn,Sn Error,
// This is data we might care about, in most cases it'll not be all.
// The fields of this struct might also change in the future since we might not care about some values anymore...
type ScanRow struct {
	Index      uint16     `csv:"Index" json:"index"`
	Reading    uint16     `csv:"Reading No" json:"reading"`
	Time       string     `csv:"Time" json:"time"`
	Type       string     `csv:"Type" json:"type"`
	Duration   float32    `csv:"Duration" json:"duration"`
	Units      string     `csv:"Units" json:"unit"`
	SigmaValue int16      `csv:"Sigma Value" json:"sigmaValue"`
	Sequence   string     `csv:"Sequence" json:"sequence"`
	User1      string     `csv:"User1" json:"user1"`
	Flags      string     `csv:"Flags" json:"flags"`
	Boat       string     `csv:"Boat" json:"boat"`
	Operator   string     `csv:"Operator" json:"operator"`
	UserLogin  string     `csv:"User Login" json:"userLogin"`
	Pb         MetalValue `csv:"Pb" json:"pb"`
	PbError    float64    `csv:"Pb Error" json:"pbError"`
	Zn         MetalValue `csv:"Zn" json:"zn"`
	ZnError    float64    `csv:"Zn Error" json:"znError"`
	Cu         MetalValue `csv:"Cu" json:"cu"`
	CuError    float64    `csv:"Cu Error" json:"cuError"`
	Sn         MetalValue `csv:"Sn" json:"sn"`
	SnError    float64    `csv:"Sn Error" json:"snError"`
}

func ParseCsv(filename string) ([]ScanRow, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var scans []ScanRow
	err = gocsv.UnmarshalBytes(bytes, &scans)
	if err != nil {
		return nil, err
	}

	return scans, nil
}

// Helper function for parsing the excel file
func marshalRow(row *[]string, header *[]string) (ScanRow, error) {
	// This is quite a hacky way of doing it, converting the header and row to a 2 row csv
	rowCsv := strings.Join(*header, ",") + "\n" + strings.Join(*row, ",")

	scan := make([]ScanRow, 0, 1)

	err := gocsv.UnmarshalString(rowCsv, &scan)
	if err != nil {
		return ScanRow{}, err
	}

	scanRow := scan[0]

	parser, err := time.Parse("1/2/06 15:04", scanRow.Time)
	if err != nil {
		return ScanRow{}, err
	}

	scanRow.Time = parser.Format("2006-01-02 15:04")

	return scanRow, nil
}

func ParseExcel(filename string) ([]ScanRow, error) {
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
	var scans []ScanRow

	for i, row := range rows[1:] { // Skip the header row
		scan, err := marshalRow(&row, &header)
		if err != nil {
			return nil, fmt.Errorf("error parsing row %d: %s", i, err)
		}

		scans = append(scans, scan)
	}

	return scans, nil
}
