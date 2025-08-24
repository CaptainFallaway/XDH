package api

type Survey struct {
	Surveyor         string `json:"surveyor"`
	Date             int64  `json:"date"`
	Location         string `json:"location"`
	WestCoastFlag    bool   `json:"westCoastFlag"`
	InstrumentSerial string `json:"instrumentSerial"`
}
