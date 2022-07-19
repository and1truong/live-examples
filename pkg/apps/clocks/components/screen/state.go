package screen

import (
	"context"
	"fmt"
	
	"github.com/jfyne/live"
	"github.com/jfyne/live/page"
	"learn/pkg/apps/clocks/components/clock"
)

func newState(c *page.Component, title string) *pageState {
	ctx := context.Background()
	
	brisbaneTz, _ := page.Init(
		ctx,
		func() (*page.Component, error) {
			return clock.NewClock("clock-1", c.Handler, c.Socket, "Australia/Brisbane")
		},
	)
	
	londonTz, _ := page.Init(
		ctx,
		func() (*page.Component, error) {
			return clock.NewClock("clock-1", c.Handler, c.Socket, "Europe/London")
		},
	)
	
	return &pageState{
		Title: title,
		Clocks: []*page.Component{
			brisbaneTz,
			londonTz,
		},
	}
}

type pageState struct {
	Title           string
	ValidationError string
	Clocks          []*page.Component
}

func onMount(title string) page.MountHandler {
	return func(_ context.Context, c *page.Component) error {
		c.State = newState(c, title)
		return nil
	}
}

func onValidate(c *page.Component) func(context.Context, live.Params) (interface{}, error) {
	return func(_ context.Context, p live.Params) (interface{}, error) {
		// Get the current page component pageState.
		state, _ := c.State.(*pageState)
		
		// Get the tz coming from the form.
		tz := p.String("tz")
		
		// Try to make a new ClockState, this will return an error if the
		// timezone is not real.
		if _, err := clock.NewClockState(tz); err != nil {
			state.ValidationError = fmt.Sprintf("Timezone %s does not exist", tz)
			return state, nil
		}
		
		// If there was no error loading the clock pageState reset the
		// validation error.
		state.ValidationError = ""
		
		return state, nil
	}
}

func onAdd(c *page.Component) func(_ context.Context, p live.Params) (interface{}, error) {
	return func(_ context.Context, p live.Params) (interface{}, error) {
		// Get the current page component pageState.
		state, _ := c.State.(*pageState)
		
		// Get the timezone sent from the form input.
		tz := p.String("tz")
		if tz == "" {
			return state, nil
		}
		
		clock, err := page.Init(
			context.Background(),
			func() (*page.Component, error) {
				return clock.NewClock(fmt.Sprintf("clock-%d", len(state.Clocks)+1), c.Handler, c.Socket, tz)
			},
		)
		
		if err != nil {
			return state, err
		}
		
		state.Clocks = append(state.Clocks, clock)
		
		return state, nil
	}
}
