package chat

import (
	"context"
	"encoding/json"
	"fmt"
	
	"github.com/jfyne/live"
)

type message struct {
	ID   string // Unique ID per message so that we can use `live-update`.
	User string
	Msg  string
}

func newMessage(data interface{}) message {
	msg := message{}
	
	if raw, err := json.Marshal(data); nil != err {
		return message{}
	} else if err := json.Unmarshal(raw, &msg); nil != err {
		return message{}
	} else {
		return msg
	}
}

type State struct {
	Messages []message
}

func newState(s live.Socket) *State {
	m, ok := s.Assigns().(*State)
	
	if !ok {
		return &State{
			Messages: []message{
				{ID: live.NewID(), User: "Room", Msg: "Welcome to chat " + live.SessionID(s.Session())},
			},
		}
	}
	
	return m
}

func onMount(ctx context.Context, socket live.Socket) (interface{}, error) {
	return newState(socket), nil
}

func onSend(ctx context.Context, socket live.Socket, params live.Params) (interface{}, error) {
	state := newState(socket)
	msg := params.String("message")
	if msg == "" {
		return state, nil
	}
	
	data := map[string]interface{}{
		"ID":   live.NewID(),
		"User": live.SessionID(socket.Session()),
		"Msg":  msg,
	}
	
	if err := socket.Broadcast("insert", data); err != nil {
		return state, fmt.Errorf("failed broadcasting new message: %w", err)
	}
	
	return state, nil
}

func onInsert(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
	state := newState(s)
	state.Messages = append(state.Messages, newMessage(data))
	
	return state, nil
}
