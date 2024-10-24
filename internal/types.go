package internal

import (
	"errors"
	"strconv"
	"time"

	"github.com/ijt/go-anytime"
)

// A grouping is a colletion of info and scans for a specific boat
type Grouping struct {
	Index        int              `json:"index"`
	BoatID       string           `json:"boatID"`
	FirstDate    Date             `json:"firstDate"`
	LastDate     Date             `json:"lastDate"`
	Unit         string           `json:"unit"`
	Scans        []ScanRow        `json:"scans"`
	InvalidScans []ScanRow        `json:"invalidScans"`
	ErrorNotes   []string         `json:"errorNotes"`
	Violations   map[string]uint8 `json:"violations"`
	Operators    []string         `json:"operators"`
}

// Index,Reading No,Time,Type,Duration,Units,Sigma Value,Sequence,User1,Flags,Boat,Operator,User Login,Pb,Pb Error,Hg,Hg Error,W,W Error,Rb,Rb Error,Se,Se Error,Zn,Zn Error,Ag,Ag Error,Bk1,Bk1 Error,Mo,Mo Error,Sr,Sr Error,As,As Error,Cu,Cu Error,Ni,Ni Error,Fe,Fe Error,Cr,Cr Error,Bk2,Bk2 Error,Bk3,Bk3 Error,Bk4,Bk4 Error,Sn,Sn Error,
// This is data we might care about, in most cases it'll not be all.
// The fields of this struct might also change in the future since we might not care about some values anymore...
type ScanRow struct {
	Index      uint16  `csv:"Index" json:"index"`
	Reading    uint16  `csv:"Reading No" json:"reading"`
	Time       Date    `csv:"Time" json:"time"`
	Type       string  `csv:"Type" json:"type"`
	Duration   float32 `csv:"Duration" json:"duration"`
	Units      string  `csv:"Units" json:"unit"`
	SigmaValue int16   `csv:"Sigma Value" json:"sigmaValue"`
	Sequence   string  `csv:"Sequence" json:"sequence"`
	User1      string  `csv:"User1" json:"user1"`
	Flags      string  `csv:"Flags" json:"flags"`
	Boat       string  `csv:"Boat" json:"boat"`
	Operator   string  `csv:"Operator" json:"operator"`
	UserLogin  string  `csv:"User Login" json:"userLogin"`

	Pb      MetalValue `csv:"Pb" json:"pb"`
	PbError float64    `csv:"Pb Error" json:"pbError"`
	Zn      MetalValue `csv:"Zn" json:"zn"`
	ZnError float64    `csv:"Zn Error" json:"znError"`
	Cu      MetalValue `csv:"Cu" json:"cu"`
	CuError float64    `csv:"Cu Error" json:"cuError"`
	Sn      MetalValue `csv:"Sn" json:"sn"`
	SnError float64    `csv:"Sn Error" json:"snError"`
}

type MetalValue struct {
	Value float64 `json:"value"`
	IsLod bool    `json:"isLod"`
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
	panic("not implemented")
}

type Date struct {
	Str  string
	Time int64
}

func parseTime(out *time.Time, val string) error {
	var (
		err1 error
		err2 error
	)

	(*out), err1 = anytime.Parse(val, time.Time{})

	if err1 != nil {
		(*out), err2 = time.Parse("1/2/06 15:04", val)

		if err2 != nil {
			return errors.Join(err1, err2)
		}
	}

	return nil
}

func (n *Date) UnmarshalCSV(val string) (err error) {
	temp := new(time.Time)

	err = parseTime(temp, val)
	if err != nil {
		return err
	}

	n.Time = temp.Unix()
	n.Str = temp.Format(TimeFormat)

	return nil
}

func (n *Date) MarshalCSV() (string, error) {
	panic("not implemented")
}
