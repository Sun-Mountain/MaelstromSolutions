package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any

		// Unmarshal the message body as an loosely-typed map.
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "broadcast"
		body["message"] = 1000

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
