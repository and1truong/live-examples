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
	clusters []*Cluster
}

type Cluster struct {
	PubSub *live.PubSub
	Nodes  []Node
}

type Node struct {
	Pattern string
	Engine  http.Handler
}

func (this *builder) AddCluster(name string, cluster *Cluster) {
	for i := range cluster.Nodes {
		cluster.PubSub.Subscribe(name, cluster.Nodes[i].Engine.(*live.HttpEngine))
	}
	
	this.clusters = append(this.clusters, cluster)
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
	
	for i := range this.clusters {
		for ii := range this.clusters[i].Nodes {
			node := this.clusters[i].Nodes[ii]
			http.Handle(node.Pattern, node.Engine)
		}
	}
	
	http.Handle("/live.js", live.Javascript{})
	http.Handle("/auto.js.map", live.JavascriptMap{})
	
	return http.ListenAndServe(address, nil)
}
