package pkg

import (
	"context"
	"net/http"
	
	"github.com/jfyne/live"
)

func NewLiveBuilder() *builder {
	return &builder{
		handlers: map[string]live.Handler{},
		engines:  map[string]http.Handler{},
	}
}

type builder struct {
	handlers map[string]live.Handler
	engines  map[string]http.Handler
}

func (this *builder) AddHandler(path string, handler live.Handler) {
	this.handlers[path] = handler
}

func (this *builder) AddEngine(path string, engine http.Handler) {
	this.engines[path] = engine
}

func (this *builder) Run(ctx context.Context, store live.HttpSessionStore, address string) error {
	for pattern, handler := range this.handlers {
		http.Handle(pattern, live.NewHttpHandler(store, handler))
	}
	
	for pattern, engine := range this.engines {
		http.Handle(pattern, engine)
	}
	
	http.Handle("/live.js", live.Javascript{})
	http.Handle("/auto.js.map", live.JavascriptMap{})
	
	return http.ListenAndServe(address, nil)
}
