package core

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
	Conn     net.Conn
	Nick     string
	Room     *Room
	Commands chan<- Command
}

func (c *Client) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')

		if err != nil {
			return
		}

		commandID, args, err := ParseMessage(msg)

		if err != nil {
			log.Print(err.Error())
			continue
		}

		command := Command{
			ID:     commandID,
			Client: c,
			Args:   args,
		}

		c.Commands <- command
	}
}

func (c *Client) SendMsg(msg string) {
	c.Conn.Write([]byte("> " + msg + "\n"))
}
