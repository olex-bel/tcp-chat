package test

import (
	"testing"

	"github.com/olex-bel/tcp-chat/pkg/core"
)

func TestServer(t *testing.T) {
	t.Run("should create server and run join command", func(t *testing.T) {
		addr := TestAddr{"127.0.0.1:8080"}
		conn := TestConn{addr: addr}
		server := core.NewServer()
		client := core.Client{
			Conn: &conn,
		}
		arg := core.CommandArg{
			Client:   &client,
			Argument: "test",
		}

		err := server.CommandRegistry.ExecuteCommand(core.CMD_JOIN, arg)

		if err != nil {
			t.Errorf("Error when MSG command executed %s.", err.Error())
		}
	})

	t.Run("should create server and run join command", func(t *testing.T) {
		addr := TestAddr{"127.0.0.1:8080"}
		conn := TestConn{addr: addr}
		server := core.NewServer()
		client := core.Client{
			Conn: &conn,
		}
		arg := core.CommandArg{
			Client:   &client,
			Argument: "test",
		}

		err := server.CommandRegistry.ExecuteCommand(core.CMD_MSG, arg)

		if err != nil {
			t.Errorf("Error when MSG command executed %s.", err.Error())
		}
	})

	t.Run("should create server and run nick command", func(t *testing.T) {
		addr := TestAddr{"127.0.0.1:8080"}
		conn := TestConn{addr: addr}
		server := core.NewServer()
		client := core.Client{
			Conn: &conn,
		}
		arg := core.CommandArg{
			Client:   &client,
			Argument: "test",
		}

		err := server.CommandRegistry.ExecuteCommand(core.CMD_NICK, arg)

		if err != nil {
			t.Errorf("Error when NICK command executed %s.", err.Error())
		}
	})

	t.Run("should create server and quick command", func(t *testing.T) {
		addr := TestAddr{"127.0.0.1:8080"}
		conn := TestConn{addr: addr}
		server := core.NewServer()
		client := core.Client{
			Conn: &conn,
		}
		arg := core.CommandArg{
			Client: &client,
		}

		err := server.CommandRegistry.ExecuteCommand(core.CMD_QUIT, arg)

		if err != nil {
			t.Errorf("Error when QUICK command executed %s.", err.Error())
		}
	})

	t.Run("should create server and rooms command", func(t *testing.T) {
		addr := TestAddr{"127.0.0.1:8080"}
		conn := TestConn{addr: addr}
		server := core.NewServer()
		client := core.Client{
			Conn: &conn,
		}
		arg := core.CommandArg{
			Client: &client,
		}

		err := server.CommandRegistry.ExecuteCommand(core.CMD_ROOMS, arg)

		if err != nil {
			t.Errorf("Error when ROOMS command executed %s.", err.Error())
		}
	})

	t.Run("should create new client", func(t *testing.T) {
		addr := TestAddr{"127.0.0.1:8080"}
		conn := TestConn{addr: addr}
		server := core.NewServer()
		client := server.NewClient(&conn)

		if client.Nick != "anonymous" {
			t.Errorf("Expected nick name %s got %s.", "anonymous", client.Nick)
		}

		if client.Conn.LocalAddr().String() != addr.addr {
			t.Errorf("Expected address %s got %s.", addr.addr, client.Conn.LocalAddr().String())
		}
	})
}
