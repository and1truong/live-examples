package clocks

import (
	"context"
	
	"github.com/jfyne/live"
	"github.com/jfyne/live/page"
	"learn/pkg/apps/clocks/components/screen"
)

func NewHandler() *live.BaseHandler {
	return live.NewHandler(
		page.WithComponentMount(onMount),
		page.WithComponentRenderer(),
	)
}

func onMount(ctx context.Context, handler live.Handler, s live.Socket) (*page.Component, error) {
	return screen.New("app", handler, s, "Clocks")
}
