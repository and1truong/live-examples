package chat

import (
	"context"
	"html/template"
	"log"
	"net/http"
	
	"github.com/jfyne/live"
)

// Not working!
func NewEngine(ctx context.Context, store live.HttpSessionStore) http.Handler {
	handler := live.NewHandler(withRender())
	handler.HandleMount(onMount)
	handler.HandleEvent("send", onSend)
	handler.HandleSelf("insert", onInsert)
	engine := live.NewHttpHandler(store, handler)
	
	return engine
}

func withRender() live.HandlerConfig {
	t, err := template.ParseFiles("pkg/apps/chat/layout.html", "pkg/apps/chat/view.html")
	if err != nil {
		log.Fatal(err)
	}
	
	return live.WithTemplateRenderer(t)
}
