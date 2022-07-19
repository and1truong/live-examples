package pagination

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/url"
	
	"github.com/jfyne/live"
)

const (
	itemsPerPage = 5
	nextPage     = "next-page"
)

func NewHandler() *live.BaseHandler {
	handler := live.NewHandler(withRenderConfig())
	handler.HandleMount(onMount)
	handler.HandleParams(paramHandler)
	handler.HandleEvent(nextPage, onNextPage)
	
	return handler
}

func withRenderConfig() live.HandlerConfig {
	t, err := template.ParseFiles("root.html", "pkg/apps/pagination/view.html")
	if err != nil {
		log.Fatal(err)
	}
	
	return live.WithTemplateRenderer(t)
}

func onMount(ctx context.Context, s live.Socket) (interface{}, error) {
	return newState(s), nil
}

// This gets called after mount and contains the URL query string values in the params map.
// This will also get called whenever the query string is changed on the page.
func paramHandler(ctx context.Context, socket live.Socket, params live.Params) (interface{}, error) {
	state := newState(socket)
	state.Page = params.Int("page")
	state.Items = queryItems(state.Page)
	return state, nil
}

func onNextPage(ctx context.Context, socket live.Socket, params live.Params) (interface{}, error) {
	page := params.Int("page")
	values := url.Values{}
	values.Add("page", fmt.Sprintf("%d", page))
	socket.PatchURL(values)
	
	return socket.Assigns(), nil
}
