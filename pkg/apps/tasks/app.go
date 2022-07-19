package tasks

import (
	"html/template"
	"log"
	
	"github.com/jfyne/live"
)

func NewHandler() *live.BaseHandler {
	handler := live.NewHandler(withRenderConfig())
	handler.HandleMount(onMount)
	handler.HandleEvent("validate", onValidate)
	handler.HandleEvent("save", onSave)
	handler.HandleEvent("complete", onComplete)
	
	return handler
}

func withRenderConfig() live.HandlerConfig {
	t, err := template.ParseFiles("root.html", "pkg/apps/tasks/view.html")
	if err != nil {
		log.Fatal(err)
	}
	
	return live.WithTemplateRenderer(t)
}
