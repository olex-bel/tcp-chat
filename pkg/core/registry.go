package core

import "errors"

var ErrorCommandExists = errors.New("command already exists")
var ErrorCommandIsNotDefined = errors.New("command doesn't exist")

type CommandArg struct {
	Client   *Client
	Argument string
	Server   *Server
}
type CommandFunc func(CommandArg) error
type CommandRegistry map[CommandID]CommandFunc

func NewCommandRegistry() CommandRegistry {
	return make(CommandRegistry)
}

func (m CommandRegistry) RegisterCommand(id CommandID, f CommandFunc) error {
	if _, exists := m[id]; exists {
		return ErrorCommandExists
	}

	m[id] = f

	return nil
}

func (m CommandRegistry) ExecuteCommand(id CommandID, arg CommandArg) error {
	if command, exists := m[id]; exists {
		return command(arg)
	}
	return ErrorCommandIsNotDefined
}
