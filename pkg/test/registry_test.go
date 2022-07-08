package test

import (
	"testing"

	"github.com/olex-bel/tcp-chat/pkg/core"
)

func TestRegistry(t *testing.T) {
	t.Run("should return error for empty command registry", func(t *testing.T) {
		registry := core.NewCommandRegistry()
		arg := core.CommandArg{}
		err := registry.ExecuteCommand(core.CMD_JOIN, arg)

		if err == nil {
			t.Error("Must return error if command registry is empty")
		}
	})

	t.Run("should register a new command callback", func(t *testing.T) {
		registry := core.NewCommandRegistry()
		err := registry.RegisterCommand(core.CMD_JOIN, func(c core.CommandArg) error {
			return nil
		})
		if err != nil {
			t.Error("Command was not registered.")
		}
	})

	t.Run("should execute registered command's callback", func(t *testing.T) {
		registry := core.NewCommandRegistry()
		arg := core.CommandArg{}
		err := registry.RegisterCommand(core.CMD_JOIN, func(c core.CommandArg) error {
			return nil
		})

		if err != nil {
			t.Error("Command was not registered.")
		}

		err = registry.ExecuteCommand(core.CMD_JOIN, arg)

		if err != nil {
			t.Error("Command was not executed.")
		}
	})
}
