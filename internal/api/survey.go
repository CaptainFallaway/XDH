package api

type Survey struct {
	Surveyor         string `json:"surveyor"`
	Date             int64  `json:"date"`
	Location         string `json:"location"`
	InstrumentSerial string `json:"instrumentSerial"`
}
