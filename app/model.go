package app

import (
	"XRF/templates"
	"math/rand"
	"strings"
)

// For binding to the frontend
type ModelInterface struct{}

func (a *ModelInterface) GetModels() string {
	sb := strings.Builder{}
	var statuses = [3]string{"danger", "warning", "success"}
	for i := 0; i < 100; i++ {
		x := rand.Intn(3)
		sb.WriteString(Render(templates.BoatModel(statuses[x])))
	}
	return sb.String()
}
