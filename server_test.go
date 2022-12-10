package main

import (
	"fmt"
	"net"
	"testing"
)

func TestHandleConnection(t *testing.T) {
	// create a server to use with the connection
	srv := newServer()

	// create a test connection
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// start a goroutine to handle the connection
	go handleConnection(conn, srv)

	// simulate a disconnect command being sent to the connection
	conn.Write([]byte("disconnect\n"))

	// check that the connection was removed from the server
	if len(srv.connections) != 0 {
		t.Errorf("Expected connection to be removed from server")
	}

	// simulate another command being sent to the connection
	conn.Write([]byte("command\n"))

	// check that the command was received by the server
	if len(srv.commands) != 1 {
		t.Errorf("Expected command to be received by server")
	}
}
