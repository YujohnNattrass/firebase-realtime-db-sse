package main

import (
	"fmt"

	"github.com/r3labs/sse"
)

func main() {
	// This one works
	client := sse.NewClient("https://solid-mountain-332320-default-rtdb.firebaseio.com/outer/small/.json")

	// This one does not work
	// client := sse.NewClient("https://solid-mountain-332320-default-rtdb.firebaseio.com/outer/large/.json")

	client.Subscribe("messages", func(msg *sse.Event) {
		fmt.Println(string(msg.Data))
	})
}
