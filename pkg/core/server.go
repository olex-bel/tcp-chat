package core

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	Rooms           map[string]*Room
	Commands        chan Command
	CommandRegistry CommandRegistry
}

func NewServer() *Server {
	commandRegistry := NewCommandRegistry()

	commandRegistry.RegisterCommand(CMD_JOIN, JoinRoomCommand)
	commandRegistry.RegisterCommand(CMD_NICK, SetNickCommand)
	commandRegistry.RegisterCommand(CMD_ROOMS, GetListRoomsCommand)
	commandRegistry.RegisterCommand(CMD_MSG, SendMessageCommand)
	commandRegistry.RegisterCommand(CMD_QUIT, QuitCommand)

	return &Server{
		Rooms:           make(map[string]*Room),
		Commands:        make(chan Command),
		CommandRegistry: commandRegistry,
	}
}

func (s *Server) Run() {
	for cmd := range s.Commands {
		arg := CommandArg{
			Client:   cmd.Client,
			Argument: cmd.Args,
			Server:   s,
		}
		s.CommandRegistry.ExecuteCommand(cmd.ID, arg)
	}
}

func (s *Server) NewClient(conn net.Conn) *Client {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())

	return &Client{
		Conn:     conn,
		Nick:     "anonymous",
		Commands: s.Commands,
	}
}

func (s *Server) QuitCurrentRoom(client *Client) {
	if client.Room != nil {
		oldRoom := s.Rooms[client.Room.Name]
		delete(s.Rooms[client.Room.Name].Members, client.Conn.RemoteAddr())
		oldRoom.Broadcast(client, fmt.Sprintf("%s has left the room", client.Nick))
	}
}

func JoinRoomCommand(arg CommandArg) error {
	client := arg.Client
	roomName := arg.Argument
	server := arg.Server

	if len(roomName) < 2 {
		client.SendMsg("room name is required. usage: /join ROOM_NAME")
		return errors.New("room argument is missing")
	}

	room, ok := server.Rooms[roomName]
	if !ok {
		room = &Room{
			Name:    roomName,
			Members: make(map[net.Addr]*Client),
		}
		server.Rooms[roomName] = room
	}
	room.Members[client.Conn.RemoteAddr()] = client

	server.QuitCurrentRoom(client)
	client.Room = room

	room.Broadcast(client, fmt.Sprintf("%s joined the room", client.Nick))

	client.SendMsg(fmt.Sprintf("welcome to %s", roomName))

	return nil
}

func SetNickCommand(arg CommandArg) error {
	client := arg.Client
	nick := arg.Argument

	if len(nick) == 0 {
		client.SendMsg("nick is required. usage: /nick NAME")
		return errors.New("nick argument is missing")
	}

	client.Nick = nick
	client.SendMsg(fmt.Sprintf("all right, I will call you %s", client.Nick))

	return nil
}

func GetListRoomsCommand(arg CommandArg) error {
	var rooms []string

	for name := range arg.Server.Rooms {
		rooms = append(rooms, name)
	}

	arg.Client.SendMsg(fmt.Sprintf("available rooms: %s", strings.Join(rooms, ", ")))

	return nil
}

func SendMessageCommand(arg CommandArg) error {
	client := arg.Client
	msg := arg.Argument

	if len(msg) == 0 {
		client.SendMsg("message is required, usage: /msg MSG")
		return errors.New("message argument is missing")
	}

	if client.Room == nil {
		client.SendMsg("please join a room")
		return errors.New("room is not selected")
	}

	client.Room.Broadcast(client, fmt.Sprintf("%s :%s", client.Nick, msg))

	return nil
}

func QuitCommand(arg CommandArg) error {
	client := arg.Client
	clientAddr := client.Conn.RemoteAddr().String()
	log.Printf("client has left the chat: %s", clientAddr)
	arg.Server.QuitCurrentRoom(client)
	client.SendMsg("bye")
	client.Conn.Close()
	return nil
}
