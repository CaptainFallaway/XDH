package internal

import (
	"errors"
	"strconv"
	"time"

	"github.com/ijt/go-anytime"
)

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
	n.Text = temp.Format(TimeFormat)

	return nil
}

func (n *Date) MarshalCSV() (string, error) {
	panic("not implemented")
}
