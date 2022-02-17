package clocks

import (
	"context"
	
	"github.com/jfyne/live"
	"github.com/jfyne/live/page"
	"learn/pkg/app/clocks/components"
)

func NewHandler() *live.BaseHandler {
	h := live.NewHandler(
		page.WithComponentMount(onMount),
		page.WithComponentRenderer(),
	)
	
	return h
}

func onMount(ctx context.Context, handler live.Handler, s live.Socket) (*page.Component, error) {
	return components.NewPage("app", handler, s, "Clocks")
}
