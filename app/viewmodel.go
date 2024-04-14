package app

import (
	"fmt"

	"github.com/CaptainFallaway/XDH/data_pipeline"
)

const FilePath string = "sample_data/Praktik 2024-03-16.xls"

type State struct {
	FilePath     string
	SortingMetal string
	Scans        []data_pipeline.ScanRow
}

// For binding to the frontend
type viewModelInterface struct {
	State *State
}

func NewViewModel() *viewModelInterface {
	return &viewModelInterface{
		State: &State{
			FilePath:     FilePath,
			SortingMetal: "Sn",
		},
	}
}

func (a *viewModelInterface) ReadFile() {
	a.State.FilePath = FilePath
	temp, err := data_pipeline.ParseExcel(a.State.FilePath)
	if err != nil {
		panic(err)
	}
	a.State.Scans = temp
}

func (a *viewModelInterface) GroupByBoatID() string {
	grouping := data_pipeline.GroupByBoat(&a.State.Scans)

	sorted := data_pipeline.SortByViolations(&grouping, "Sn")

	fmt.Println(sorted)

	return ""
}
