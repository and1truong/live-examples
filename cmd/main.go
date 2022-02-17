package main

import (
	"context"
	"fmt"
	
	"github.com/jfyne/live"
	"learn/pkg"
	"learn/pkg/app/chart"
	"learn/pkg/app/chat"
	"learn/pkg/app/clocks"
	"learn/pkg/app/counter"
)

func cookieStore() live.HttpSessionStore {
	return live.NewCookieStore("session-name", []byte("weak-secret"))
}

func main() {
	ctx := context.Background()
	store := cookieStore()
	builder := pkg.NewLiveBuilder()
	builder.AddHandler("/counter", counter.NewHandler())
	builder.AddHandler("/clocks", clocks.NewHandler())
	builder.AddEngine("/charts", chart.NewEngine(ctx, store))
	builder.AddEngine("/chat", chat.NewEngine(ctx, store))
	
	if err := builder.Run(ctx, store, ":8181"); nil != err {
		fmt.Println("server error: ", err.Error())
	}
}