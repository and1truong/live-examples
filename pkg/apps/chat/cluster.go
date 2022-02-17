package chat

import (
	"context"
	"log"

	"github.com/jfyne/live"
	_ "gocloud.dev/pubsub/mempubsub"
	"learn/pkg"
)

func NewCluster(ctx context.Context, store live.HttpSessionStore) *pkg.Cluster {
	node1 := NewEngine(ctx, store).(*live.HttpEngine)
	node2 := NewEngine(ctx, store).(*live.HttpEngine)

	ps, err := pkg.NewPubSub(ctx, true)
	if nil != err {
		log.Fatal(err)
	}

	return &pkg.Cluster{
		PubSub: ps,
		Nodes: []pkg.Node{
			{"/chat/one", node1},
			{"/chat/two", node2},
		},
	}
}
