package core

import (
	"fmt"
	"strings"
)

type CommandID int

const (
	CMD_NICK CommandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

var CommandMap = map[string]CommandID{
	"/join":  CMD_JOIN,
	"/msg":   CMD_MSG,
	"/nick":  CMD_NICK,
	"/quit":  CMD_QUIT,
	"/rooms": CMD_ROOMS,
}

type Command struct {
	ID     CommandID
	Client *Client
	Args   string
}

func ParseMessage(msg string) (CommandID, string, error) {
	msg = strings.Trim(msg, "\r\n")
	fields := strings.SplitN(msg, " ", 2)

	commandName := strings.TrimSpace(fields[0])
	args := ""

	if len(fields) > 1 {
		args = fields[1]
	}

	commandID, exists := CommandMap[commandName]

	if !exists {
		return -1, "", fmt.Errorf("unknown command: %s", commandName)
	}

	return commandID, args, nil
}
