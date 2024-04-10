package app

import (
	"github.com/CaptainFallaway/XDH/data_handling"
	"github.com/CaptainFallaway/XDH/templates"
)

const FilePath string = "sample_data/Praktik 2024-03-16.csv"

// For binding to the frontend
type ModelInterface struct {
	FilePath string
}

func (a *ModelInterface) GetModels() string {
	a.FilePath = FilePath
	scans, err := data_handling.ParseCsv(a.FilePath)
	if err != nil {
		panic(err)
	}

	group := data_handling.GroupByBoats(scans)

	var sorted []data_handling.Scan
	for boat, scans := range group {

	}

	return Render(templates.BoatModel(scans))
}

func (a *ModelInterface) GetModel(id string, active bool) string {
	return Render(templates.BoatModelContent(active))
}
