// Domain specific types
// These types are both accessed here in the go code
// And the frontend code

package internal

// Grouping is a collection of info and scans for a specific boat
type Grouping struct {
	Index        int              `json:"index"` // An index to help with the sorting, represent on a abstract basis what rows where read
	BoatID       string           `json:"boatID"`
	FirstDate    int64            `json:"firstDate"`
	LastDate     int64            `json:"lastDate"`
	Unit         string           `json:"unit"`         // The unit that is used for the metal values
	Scans        []Scan           `json:"scans"`        // Valid scans that fall above the minimum scan time
	InvalidScans []Scan           `json:"invalidScans"` // Scans that fall under the minimum scan time
	ErrorNotes   []string         `json:"errorNotes"`   // Notes that are accumulator during build of grouping, like if there is less than 8 scans
	Violations   map[string]uint8 `json:"violations"`   // A violation count map, it is used for sorting based on metal violations
	Operators    []string         `json:"operators"`    // The operators that were found in the scans for this boat
}

type Scan struct {
	Reading    int             `json:"reading"`
	Duration   float64         `json:"duration"`
	Operator   string          `json:"operator"`
	Date       int64           `json:"date"`
	Pb         float64         `json:"pb"`
	Zn         float64         `json:"zn"`
	Cu         float64         `json:"cu"`
	Sn         float64         `json:"sn"`
	Violations map[string]bool `json:"violations"`
}

type SessionInfo struct {
	Uid              string `json:"uid"`
	Surveyor         string `json:"surveyor"`
	Date             int64  `json:"date"`
	Location         string `json:"location"`
	InstrumentSerial string `json:"instrumentSerial"`
}

type Session struct {
	Session   *SessionInfo
	Groupings []Grouping
}

func NewSession(session *SessionInfo, groupings []Grouping) *Session {
	return &Session{
		Session:   session,
		Groupings: groupings,
	}
}
