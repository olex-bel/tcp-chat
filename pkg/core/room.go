package core

import "net"

type Room struct {
	Name    string
	Members map[net.Addr]*Client
}

func (r *Room) Broadcast(c *Client, msg string) error {
	for _, client := range r.Members {
		if client != c {
			client.SendMsg(msg)
		}
	}

	return nil
}
