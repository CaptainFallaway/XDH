package data_handling

type Sorted struct {
	BoatID   string
	Operator string
	Scans    []Scan
}

func SortByCriticality(scans []Scan) []Sorted {
	var ret []Sorted = make([]Sorted, len(scans))

	for i, scan := range scans {
		
}
