package counter

import (
	"context"
	"html/template"
	"log"

	"github.com/jfyne/live"
)

func NewHandler() *live.BaseHandler {
	h := live.NewHandler(withRender())
	h.HandleMount(onMount)
	h.HandleEvent("inc", onIncrease)
	h.HandleEvent("dec", onDecrease)

	return h
}

func withRender() live.HandlerConfig {
	t, err := template.ParseFiles("root.html", "pkg/apps/counter/view.html")
	if err != nil {
		log.Fatal(err)
	}

	return live.WithTemplateRenderer(t)
}

func onMount(ctx context.Context, s live.Socket) (interface{}, error) {
	c := newAppState(s)
	c.Value = 1000

	return c, nil
}
