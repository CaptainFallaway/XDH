package data_handling

import "slices"

type Grouping interface {
	Keys() []string
	Get(key string) []Scan
}

type grouping struct {
	data map[string][]Scan
}

func (g grouping) Keys() []string {
	keys := make([]string, 0, len(g.data))

	for key := range g.data {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	return keys
}

func (g grouping) Get(key string) []Scan {
	return g.data[key]
}

func GroupByBoats(scans []Scan) Grouping {
	boats := make(map[string][]Scan)

	for _, scan := range scans {
		boats[scan.Boat] = append(boats[scan.Boat], scan)
	}

	return grouping{boats}
}

func GroupByOperator(scans []Scan) Grouping {
	operators := make(map[string][]Scan)

	for _, scan := range scans {
		operators[scan.Operator] = append(operators[scan.Operator], scan)
	}

	return grouping{operators}
}
