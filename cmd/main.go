package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jfyne/live"
	"github.com/jfyne/live-contrib/livefiber"
	"learn/pkg"
	"learn/pkg/apps/chart"
	"learn/pkg/apps/chat"
	"learn/pkg/apps/clock"
	"learn/pkg/apps/clocks"
	"learn/pkg/apps/counter"
	"learn/pkg/apps/pagination"
	"learn/pkg/apps/task"
)

func cookieStore() live.HttpSessionStore {
	return live.NewCookieStore("session-name", []byte("weak-secret"))
}

func main() {
	if os.Getenv("MODE") == "fiber" {
		fiberMain()
	} else {
		builderMain()
	}
}

func builderMain() {
	ctx := context.Background()
	store := cookieStore()

	builder := pkg.NewLiveBuilder()
	builder.AddHandler("/counter", counter.NewHandler())
	builder.AddHandler("/clocks", clocks.NewHandler())
	builder.AddHandler("/clock", clock.NewHandler())
	builder.AddEngine("/charts", chart.NewEngine(ctx, store))
	builder.AddEngine("/chat", chat.NewEngine(ctx, store))
	builder.AddCluster("chat-app", chat.NewCluster(ctx, store))
	builder.AddHandler("/tasks", task.NewHandler())
	builder.AddHandler("/pagination", pagination.NewHandler())

	if err := builder.Run(ctx, store, ":8181"); nil != err {
		fmt.Println("server error: ", err.Error())
	}
}

func fiberMain() {
	store := session.New()
	handler := counter.NewHandler()
	engine := livefiber.NewHandler(store, handler)
	handlers := engine.Handlers()

	app := fiber.New()
	app.Get("/", handlers...)
	app.Get("/live.js", adaptor.HTTPHandler(live.Javascript{}))
	app.Get("/auto.js.map", adaptor.HTTPHandler(live.JavascriptMap{}))
	log.Fatal(app.Listen(":8181"))
}
