package parsers

import (
	"errors"
	"strconv"
	"time"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/ijt/go-anytime"
)

// ScanRow These are the values we care about in the csv and excel files, most of them are only here for future implementations
type ScanRow struct {
	Index      int     `csv:"Index"`
	Reading    int     `csv:"Reading No"`
	Time       Date    `csv:"Time"`
	Type       string  `csv:"Type"`
	Duration   float64 `csv:"Duration"`
	Units      string  `csv:"Units"`
	SigmaValue int16   `csv:"Sigma Value"`
	Sequence   string  `csv:"Sequence"`
	User1      string  `csv:"User1"`
	Flags      string  `csv:"Flags"`
	Boat       string  `csv:"Boat"`
	Operator   string  `csv:"Operator"`
	UserLogin  string  `csv:"User Login"`

	Pb      MetalValue `csv:"Pb"`
	PbError float64    `csv:"Pb Error"`
	Zn      MetalValue `csv:"Zn"`
	ZnError float64    `csv:"Zn Error"`
	Cu      MetalValue `csv:"Cu"`
	CuError float64    `csv:"Cu Error"`
	Sn      MetalValue `csv:"Sn"`
	SnError float64    `csv:"Sn Error"`
}

// There are Unmarshal and Marshal functions in typesimpl.go file
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
	panic("not expected to call marshalCSV function")
}

type Date struct {
	Text string
	Unix int64
}

func parseTime(out *time.Time, val string) error {
	var (
		err1 error
		err2 error
	)

	*out, err1 = anytime.Parse(val, time.Time{})

	if err1 != nil {
		*out, err2 = time.Parse("1/2/06 15:04", val)

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

	n.Unix = temp.Unix()
	n.Text = temp.Format(internal.TimeFormat)

	return nil
}

func (n *Date) MarshalCSV() (string, error) {
	panic("not expected to call marshalCSV function")
}
