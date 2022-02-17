package chart

import (
	"context"
	"math/rand"
	
	"github.com/jfyne/live"
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

func onMount(ctx context.Context, s live.Socket) (interface{}, error) {
	return newState(s), nil
}

func onRegenerate(ctx context.Context, socket live.Socket, d interface{}) (interface{}, error) {
	c := newState(socket)
	c.Sales = d.([]int)
	return c, nil
}
