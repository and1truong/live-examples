package chart

import (
	"context"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
	
	"github.com/jfyne/live"
)

func NewEngine(ctx context.Context, store live.HttpSessionStore) http.Handler {
	h := live.NewHandler(withRenderConfig())
	h.HandleMount(onMount)
	h.HandleSelf("regenerate", onRegenerate)
	engine := live.NewHttpHandler(store, h)
	go tick(ctx, engine)
	
	return engine
}

func tick(ctx context.Context, engine *live.HttpEngine) {
	ticker := time.NewTicker(333 * time.Millisecond)
	
	for {
		select {
		case <-ticker.C:
			_ = engine.Broadcast("regenerate", rand.Perm(9))
		
		case <-ctx.Done():
			return
		}
	}
}

func withRenderConfig() live.HandlerConfig {
	t, err := template.ParseFiles("root.html", "pkg/apps/chart/view.html")
	if err != nil {
		log.Fatal(err)
	}
	
	return live.WithTemplateRenderer(t)
}
