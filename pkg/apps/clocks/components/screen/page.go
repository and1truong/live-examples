package screen

import (
	"fmt"
	"io"
	
	"github.com/jfyne/live"
	"github.com/jfyne/live/page"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

const (
	validateTZ = "validate-tz"
	addTime    = "add-time"
)

func New(ID string, handler live.Handler, socket live.Socket, title string) (*page.Component, error) {
	return page.NewComponent(ID, handler, socket,
		page.WithRegister(
			func(c *page.Component) error {
				c.HandleEvent(validateTZ, onValidate(c))
				c.HandleEvent(addTime, onAdd(c))
				return nil
			},
		),
		page.WithMount(onMount(title)),
		page.WithRender(renderPage),
	)
}

func renderPage(w io.Writer, node *page.Component) error {
	state, ok := node.State.(*pageState)
	if !ok {
		return fmt.Errorf("could not get pageState")
	}
	
	props := c.HTML5Props{
		Title:    state.Title,
		Language: "en",
		Head: []g.Node{
			h.StyleEl(h.Type("text/css"),
				g.Raw(`body {font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol"; }`),
			),
		},
		Body: []g.Node{
			h.H1(g.Text("World Clocks")),
			h.FormEl(
				h.ID("tz-form"),
				g.Attr("live-change", node.Event(validateTZ)), // c.Event scopes the events to this component.
				g.Attr("live-submit", node.Event(addTime)),
				h.Div(
					h.P(g.Text("Try Europe/London or America/New_York")),
					h.Input(h.Name("tz")),
					g.If(state.ValidationError != "", h.Span(g.Text(state.ValidationError))),
				),
				h.Input(h.Type("submit"), g.If(state.ValidationError != "", h.Disabled())),
			),
			h.Div(
				g.Group(g.Map(len(state.Clocks), func(idx int) g.Node {
					return page.Render(state.Clocks[idx])
				})),
			),
			h.Script(h.Src("/live.js")),
		},
	}
	
	return c.HTML5(props).Render(w)
}
