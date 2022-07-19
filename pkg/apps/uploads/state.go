package uploads

import (
	"context"
	
	"github.com/jfyne/live"
)

type state struct {
	Uploads []string
}

func newState(s live.Socket) *state {
	m, ok := s.Assigns().(*state)
	if !ok {
		return &state{
			Uploads: []string{},
		}
	}
	
	return m
}

func onMount(ctx context.Context, s live.Socket) (interface{}, error) {
	s.AllowUploads(&live.UploadConfig{
		Name:     "photos",        // Name refers to the name of the file input field.
		MaxFiles: 3,               // We are accepting a maximum of 3 files.
		MaxSize:  1 * 1024 * 1024, // For each of those files we are only allowing them to be 1MB.
		Accept:   []string{"image/png", "image/jpeg"},
	})
	
	return newState(s), nil
}
