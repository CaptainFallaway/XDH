package internal

type metalPolicy struct {
	AmmountOfMetals uint8
	Metals          []string
	SnViolation     float64
	CuViolation     float64
	ZnViolation     float64
	PbViolation     float64
}

var MetalPolicy metalPolicy = metalPolicy{
	AmmountOfMetals: 4,
	Metals:          []string{"Pb", "Zn", "Cu", "Sn"}, // Theses are the metals that we care about in this project
	PbViolation:     100,
	ZnViolation:     100,
	CuViolation:     1000,
	SnViolation:     100,
}

const (
	TimeFormat               string  = "2006-01-02 15:04" // Standardised time format
	ValidMinimumScanTime     float32 = 20                 // Minimum time for a scan to be valid
	ExpectedAmmountOfColumns int     = 21                 // Expected ammoud of colums when parsing either csv or excel files
)
