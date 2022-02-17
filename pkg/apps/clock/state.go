package clock

import (
	"context"
	"time"

	"github.com/jfyne/live"
)

const (
	tick = "tick"
)

type State struct {
	Time time.Time
}

func newClock(s live.Socket) *State {
	c, ok := s.Assigns().(*State)
	if !ok {
		return &State{
			Time: time.Now(),
		}
	}
	return c
}

func (c State) FormattedTime() string {
	return c.Time.Format("15:04:05")
}

func onMount(ctx context.Context, s live.Socket) (interface{}, error) {
	c := newClock(s)

	if s.Connected() {
		go func() {
			time.Sleep(1 * time.Second)
			s.Self(ctx, tick, time.Now())
		}()
	}

	return c, nil
}

func onTick(ctx context.Context, s live.Socket, d interface{}) (interface{}, error) {
	// Get our model
	c := newClock(s)
	// Update the time.
	c.Time = d.(time.Time)
	// Send ourselves another tick in a second.
	go func(sock live.Socket) {
		time.Sleep(1 * time.Second)
		sock.Self(ctx, tick, time.Now())
	}(s)
	return c, nil
}
