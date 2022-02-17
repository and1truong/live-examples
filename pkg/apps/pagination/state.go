package pagination

import (
	"github.com/jfyne/live"
)

type State struct {
	Items []string
	Page  int
}

func newState(s live.Socket) *State {
	newState, ok := s.Assigns().(*State)
	if !ok {
		newState = &State{
			Items: []string{},
			Page:  0,
		}
	}
	return newState
}

func (s State) NextPage() int {
	return s.Page + 1
}
