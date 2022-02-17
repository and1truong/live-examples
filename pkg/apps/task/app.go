package task

import (
	"html/template"
	"log"

	"github.com/jfyne/live"
)

func newState(s live.Socket) *state {
	m, ok := s.Assigns().(*state)
	if !ok {
		return &state{
			Form: form{
				Errors: map[string]string{},
			},
		}
	}

	m.Form.Errors = map[string]string{}
	return m
}

func NewHandler() *live.BaseHandler {
	handler := live.NewHandler(withRender())
	handler.HandleMount(onMount)
	handler.HandleEvent("validate", onValidate)
	handler.HandleEvent("save", onSave)
	handler.HandleEvent("complete", onComplete)

	return handler
}

func withRender() live.HandlerConfig {
	t, err := template.ParseFiles("root.html", "pkg/apps/task/view.html")
	if err != nil {
		log.Fatal(err)
	}

	return live.WithTemplateRenderer(t)
}
