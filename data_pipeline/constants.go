package data_pipeline

type metalPolicy struct {
	AmmountOfMetals uint8
	Metals          [4]string
	PbViolation     float64
	ZnViolation     float64
	CuViolation     float64
	SnViolation     float64
}

var MetalPolicy metalPolicy = metalPolicy{
	AmmountOfMetals: 4,
	Metals:          [4]string{"Pb", "Zn", "Cu", "Sn"}, // Theses are the metals that we care about in this project
	PbViolation:     100,
	ZnViolation:     100,
	CuViolation:     100,
	SnViolation:     100,
}

const (
	TimeFormat           string  = "2006-01-02 15:04" // Global time format used in the project, the csv and xls files have different formats
	ValidMinimumScanTime float32 = 20                 // Minimum time for a scan to be valid
)
