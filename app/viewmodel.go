package app

import (
	"strconv"

	"github.com/CaptainFallaway/XDH/data_pipeline"
	"github.com/CaptainFallaway/XDH/templates"
)

func NewViewModelInterface() *ViewModelInterface {
	return &ViewModelInterface{
		FilePath:     "sample_data/Stockholms SS omgång 4 + Ängsvik BK.xls",
		SortingMetal: "Sn",
	}
}

// For binding to the frontend
type ViewModelInterface struct {
	FilePath     string
	SortingMetal string
	CurrGrouping string
	Scans        []data_pipeline.ScanRow
}

func (vm *ViewModelInterface) ReadFile() {
	temp, err := data_pipeline.ParseExcel(vm.FilePath)
	if err != nil {
		panic(err)
	}
	vm.Scans = temp
}

func (vm *ViewModelInterface) GetBoatIDModel(index string, expanded string) string {
	convIndex, err := strconv.ParseUint(index, 10, 32)
	if err != nil {
		panic(err)
	}

	convExpanded, err := strconv.ParseBool(expanded)
	if err != nil {
		panic(err)
	}

	grouping := data_pipeline.GroupByBoat(&vm.Scans)

	for _, group := range grouping {
		if group.Index == uint32(convIndex) {
			return Render(templates.BoatIDModelContent(group, !convExpanded))
		}
	}

	return "Something went horribly wrong!"
}

func (vm *ViewModelInterface) GroupByBoatID() string {
	vm.CurrGrouping = "BoatID"

	grouping := data_pipeline.GroupByBoat(&vm.Scans)

	data_pipeline.SortByViolations(&grouping, vm.SortingMetal)

	return Render(templates.BoatIDModelList(&grouping, vm.SortingMetal))
}

func (vm *ViewModelInterface) GroupByOperator() string {
	vm.CurrGrouping = "Operator"

	grouping := data_pipeline.GroupByOperator(&vm.Scans)

	data_pipeline.SortByViolations(&grouping, vm.SortingMetal)

	return Render(templates.OperatorModelList(&grouping, vm.SortingMetal))
}

func (vm *ViewModelInterface) SetSortingMetal(metal string) string {
	if MetalInMetalPolicy(metal) {
		vm.SortingMetal = metal
	}

	if vm.CurrGrouping == "BoatID" {
		return vm.GroupByBoatID()
	} else {
		return vm.GroupByOperator()
	}
}
