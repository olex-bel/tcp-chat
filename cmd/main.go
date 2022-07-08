package main

import (
	"log"
	"net"

	"github.com/olex-bel/tcp-chat/pkg/core"
)

func main() {
	server := core.NewServer()

	go server.Run()

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on :8888")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		client := server.NewClient(conn)
		go client.ReadInput()
	}
}
