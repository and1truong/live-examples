package screen

import (
	"fmt"
	
	"github.com/jfyne/live/page"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

type Render struct {
	c *page.Component
}

type nodeHolder func(children ...g.Node) g.Node
type nodeRender func(c *page.Component) (g.Node, error)

type Outcome struct {
	node g.Node
	err  error
}

var stateError = fmt.Errorf("could not get state")

func (this *Render) Div(children ...any) Outcome {
	return this.el(h.Div, children...)
}

func (this *Render) P(children ...any) Outcome {
	return this.el(h.P, children...)
}

func (this *Render) H1(children ...any) Outcome {
	return this.el(h.H1, children...)
}

func (this *Render) Form(children ...any) Outcome {
	return this.el(h.FormEl, children...)
}

func (this *Render) Script(children ...any) Outcome {
	return this.el(h.Script, children...)
}

func (this *Render) input(children ...any) Outcome {
	return this.el(h.Input, children...)
}

func (this *Render) onChange(event string) g.Node {
	return g.Attr("live-change", this.c.Event(event))
}

func (this *Render) onSubmit(event string) g.Node {
	return g.Attr("live-submit", this.c.Event(event))
}

func (this *Render) el(holder nodeHolder, children ...any) Outcome {
	elements := []g.Node{}
	for i := range children {
		switch child := children[i].(type) {
		case string:
			elements = append(elements, g.Text(child))
		
		case g.Node:
			elements = append(elements, child)
		
		case nodeRender:
			element, err := child(this.c)
			if nil != err {
				return Outcome{err: err}
			} else {
				elements = append(elements, element)
			}
		}
	}
	
	return Outcome{
		node: holder(elements...),
		err:  nil,
	}
}
