package test

import (
	"testing"

	"github.com/olex-bel/tcp-chat/pkg/core"
)

func TestParseMessage(t *testing.T) {
	t.Run("should return error if invalid format", func(t *testing.T) {
		_, _, err := core.ParseMessage("")

		if err == nil {
			t.Error("Should return error for invalid message")
		}
	})

	t.Run("should return error if command is invalid", func(t *testing.T) {
		_, _, err := core.ParseMessage("/test")

		if err == nil {
			t.Error("Should return error if command is invalid")
		}
	})

	t.Run("should parse join command", func(t *testing.T) {
		msg := "/join test"
		id, args, err := core.ParseMessage(msg)

		if err != nil {
			t.Errorf("Cannot parse join command: %s\n", msg)
		}

		if id != core.CMD_JOIN {
			t.Errorf("Expected ID %d got %d", core.CMD_JOIN, id)
		}

		if len(args) == 0 {
			t.Error("Expected 1 argument")
		}
	})

	t.Run("should parse quit command", func(t *testing.T) {
		msg := "/quit "
		id, _, err := core.ParseMessage(msg)

		if err != nil {
			t.Errorf("Cannot parse quit command: %s\n", msg)
		}

		if id != core.CMD_QUIT {
			t.Errorf("Expected ID %d got %d", core.CMD_QUIT, id)
		}
	})

	t.Run("should parse nick command", func(t *testing.T) {
		msg := "/nick test"
		id, args, err := core.ParseMessage(msg)

		if err != nil {
			t.Errorf("Cannot parse nick command: %s\n", msg)
		}

		if id != core.CMD_NICK {
			t.Errorf("Expected ID %d got %d", core.CMD_NICK, id)
		}

		if len(args) == 0 {
			t.Error("Expected 1 argument")
		}
	})

	t.Run("should parse msg command", func(t *testing.T) {
		msg := "/msg test"
		id, args, err := core.ParseMessage(msg)

		if err != nil {
			t.Errorf("Cannot parse msg command: %s\n", msg)
		}

		if id != core.CMD_MSG {
			t.Errorf("Expected ID %d got %d", core.CMD_MSG, id)
		}

		if len(args) == 0 {
			t.Error("Expected 1 argument")
		}
	})

	t.Run("should parse rooms command", func(t *testing.T) {
		msg := "/rooms"
		id, _, err := core.ParseMessage(msg)

		if err != nil {
			t.Errorf("Cannot parse rooms command: %s\n", msg)
		}

		if id != core.CMD_ROOMS {
			t.Errorf("Expected ID %d got %d", core.CMD_ROOMS, id)
		}
	})
}
