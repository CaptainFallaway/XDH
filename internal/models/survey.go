package models

type Survey struct {
	Uid              string   `json:"uid"`
	Surveyor         string   `json:"surveyor"`
	Date             int64    `json:"date"`
	Location         string   `json:"location"`
	InstrumentSerial string   `json:"instrumentSerial"`
	GroupingIds      []string `json:"groupingIds"`
}
