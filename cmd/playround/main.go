package main

import (
	"encoding/json"
	"fmt"

	"github.com/jfyne/live"
)

func main() {
	msg := live.Event{
		T:    "insert",
		ID:   12345,
		Data: nil,
		SelfData: map[string]interface{}{
			"ID":   "c86sqi7m20rk3r8ni6ag",
			"User": "c86q387m20riakbgvejg",
			"Msg":  "Ping",
		},
	}

	data, _ := json.Marshal(live.TransportMessage{Topic: "my-app", Msg: msg})

	fmt.Println("encoded: ", string(data))
}
