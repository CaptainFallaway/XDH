package app

import (
	"context"

	"github.com/CaptainFallaway/XDH/data_pipeline"
	"github.com/a-h/templ"
)

type Component struct {
	Html string
}

// Write is a method that satisfies the io.Writer interface
func (c *Component) Write(p []byte) (int, error) {
	c.Html = string(p)
	return len(p), nil
}

// You should only care about this function for rendering the templates
// Returns the rendered HTML as a string for the swapping in the frontend
func Render(comp templ.Component) string {
	var component Component
	comp.Render(context.Background(), &component)
	return component.Html
}

func MetalInMetalPolicy(metal string) bool {
	for _, m := range data_pipeline.MetalPolicy.Metals {
		if m == metal {
			return true
		}
	}
	return false
}
