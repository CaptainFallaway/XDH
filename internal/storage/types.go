package storage

type Survey struct {
	Uid              string ` gorm:"primaryKey"`
	Surveyor         string
	Date             int64
	Location         string
	InstrumentSerial string
}

type Grouping struct {
	Uid        string `gorm:"primaryKey"`
	Index      int
	BoatID     string
	FirstDate  int64
	LastDate   int64
	Unit       string
	ErrorNotes []string         `gorm:"type:json"`
	Violations map[string]uint8 `gorm:"type:json"`
	Operators  []string         `gorm:"type:json"`

	SurveyID string `gorm:"index"`
}

type Scan struct {
	Uid        string `gorm:"primaryKey"`
	Reading    int
	Duration   float64
	Operator   string
	Date       int64
	Pb         float64
	Zn         float64
	Cu         float64
	Sn         float64
	Violations map[string]bool

	GroupingID string `gorm:"index"`
}
