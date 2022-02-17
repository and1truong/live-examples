package counter

import (
	"context"

	"github.com/jfyne/live"
)

func newAppState(socket live.Socket) *appState {
	state, ok := socket.Assigns().(*appState)
	if !ok {
		return &appState{}
	}
	return state
}

type appState struct {
	Value int
}

func onIncrease(ctx context.Context, socket live.Socket, _ live.Params) (interface{}, error) {
	state := newAppState(socket)
	state.Value += 1
	return state, nil
}

func onDecrease(ctx context.Context, socket live.Socket, _ live.Params) (interface{}, error) {
	state := newAppState(socket)
	state.Value -= 1
	return state, nil
}
