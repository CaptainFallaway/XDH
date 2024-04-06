package app

import (
	"XRF/templates"
	"context"
	"fmt"
	"strings"

	"github.com/a-h/templ"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type Component struct {
	Html string
}

// Write is a method that satisfies the io.Writer interface
func (c *Component) Write(p []byte) (int, error) {
	c.Html = string(p)

	return len(p), nil
}

func render(comp templ.Component) string {
	var component Component
	comp.Render(context.Background(), &component)
	return component.Html

}

func (a *App) GetHello() string {
	sb := strings.Builder{}
	for i := 0; i < 1000; i++ {
		sb.WriteString(render(templates.Hello()))
	}
	return sb.String()
}
