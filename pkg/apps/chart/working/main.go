package main

import (
	"context"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
	
	"github.com/jfyne/live"
	"learn/pkg"
)

type state struct {
	Sales []int
}

func newState(s live.Socket) *state {
	d, ok := s.Assigns().(*state)
	if !ok {
		return &state{
			Sales: rand.Perm(9),
		}
	}
	return d
}

func cookieStore() live.HttpSessionStore {
	return live.NewCookieStore("session-name", []byte("weak-secret"))
}

func main() {
	store := cookieStore()
	ctx := context.Background()
	app := getApp()
	engine := live.NewHttpHandler(store, app)
	go tick(ctx, engine)
	
	http.Handle("/chart", engine)
	builder := pkg.NewLiveBuilder()
	builder.AddEngine("/wip", engine)
	builder.Run(ctx, store, ":8181")
}

func getApp() *live.BaseHandler {
	t, err := template.ParseFiles("root.html", "pkg/apps/chart/view.html")
	if err != nil {
		log.Fatal(err)
	}
	
	app := live.NewHandler(live.WithTemplateRenderer(t))
	app.HandleMount(onMount)
	app.HandleSelf("regenerate", onRegenerate)
	
	return app
}

func onMount(ctx context.Context, s live.Socket) (interface{}, error) {
	return newState(s), nil
}

func tick(ctx context.Context, e *live.HttpEngine) {
	rand.Seed(time.Now().Unix())
	ticker := time.NewTicker(2 * time.Second)
	for {
		<-ticker.C
		_ = e.Broadcast("regenerate", rand.Perm(9))
	}
}

func onRegenerate(ctx context.Context, s live.Socket, d interface{}) (interface{}, error) {
	c := newState(s)
	c.Sales = d.([]int)
	return c, nil
}
