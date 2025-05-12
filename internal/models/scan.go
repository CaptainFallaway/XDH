package models

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
