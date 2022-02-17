package main

import (
	"context"
	"fmt"
	
	"github.com/jfyne/live"
	"learn/pkg"
	"learn/pkg/apps/chart"
	"learn/pkg/apps/chat"
	"learn/pkg/apps/clocks"
	"learn/pkg/apps/counter"
	"learn/pkg/apps/task"
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
	builder.AddCluster("chat-app", chat.NewCluster(ctx, store))
	builder.AddHandler("/tasks", task.NewHandler())
	
	if err := builder.Run(ctx, store, ":8181"); nil != err {
		fmt.Println("server error: ", err.Error())
	}
}
