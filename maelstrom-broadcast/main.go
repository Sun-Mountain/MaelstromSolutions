package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	values := []float64{}

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any

		// Unmarshal the message body as an loosely-typed map.
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		reply := map[string]string{"type": "broadcast_ok"}

		newValue := body["message"].(float64)

		values = append(values, newValue)

		return n.Reply(msg, reply)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any

		// Unmarshal the message body as an loosely-typed map.
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		reply := map[string]string{"type": "topology_ok"}

		return n.Reply(msg, reply)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any

		// Unmarshal the message body as an loosely-typed map.
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		reply := map[string]any{"type": "read_ok", "messages": values}

		return n.Reply(msg, reply)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
