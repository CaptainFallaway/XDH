package internal

// Grouping is a collection of info and scans for a specific boat
type Grouping struct {
	Index        int              `json:"index"` // A index to help with the sorting, represent on a abstract basis what rows where read
	BoatID       string           `json:"boatID"`
	FirstDate    Date             `json:"firstDate"`
	LastDate     Date             `json:"lastDate"`
	Unit         string           `json:"unit"`         // The unit that is used for the metal values
	Scans        []ScanRow        `json:"scans"`        // Valid scans that fall above the minimum scan time
	InvalidScans []ScanRow        `json:"invalidScans"` // Scans that fall under the minimum scan time
	ErrorNotes   []string         `json:"errorNotes"`   // Notes that are accumulator during build of grouping, like if there is less than 8 scans
	Violations   map[string]uint8 `json:"violations"`   // A violation count map, it is used for sorting based on metal violations
	Operators    []string         `json:"operators"`    // The operators that were found in the scans for this boat
}

// ScanRow These are the values we care about in the csv and excel files, most of them are only here for future implementations
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

// There are Unmarshal and Marshal functions in impl file

type MetalValue struct {
	Value float64 `json:"value"`
	IsLod bool    `json:"isLod"`
}

type Date struct {
	Text string `json:"text"`
	Unix int64  `json:"unix"`
}
