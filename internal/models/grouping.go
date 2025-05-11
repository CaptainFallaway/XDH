package models

// Grouping is a collection of info for a specific boat
type Grouping struct {
	Uid        string           `json:"uid"`
	Index      int              `json:"index"` // An index to help with the sorting, represent on a abstract basis what rows where read
	BoatID     string           `json:"boatID"`
	FirstDate  int64            `json:"firstDate"`
	LastDate   int64            `json:"lastDate"`
	Unit       string           `json:"unit"`       // The unit that is used for the metal values
	ErrorNotes []string         `json:"errorNotes"` // Notes that are accumulated during build of grouping. Like if there is less than 8 scans
	Violations map[string]uint8 `json:"violations"` // A violation count map, it is used for sorting based on metal violations
	Operators  []string         `json:"operators"`  // The operators that were found in the scans for this boat
}
