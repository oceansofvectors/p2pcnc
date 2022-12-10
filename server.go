package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
)

// create a struct to hold our server's state
type server struct {
	connections map[string]net.Conn // map of client connections keyed by IP address
	commands    chan string         // channel for receiving client commands
	addconn     chan net.Conn       // channel for adding new client connections
	rmconn      chan net.Conn       // channel for removing client connections
}

// initialize a new server instance
func newServer() *server {
	return &server{
		connections: make(map[string]net.Conn),
		commands:    make(chan string),
		addconn:     make(chan net.Conn),
		rmconn:      make(chan net.Conn),
	}
}

func handleConnection(conn net.Conn, srv *server) {
	// create a scanner to read from the connection
	scanner := bufio.NewScanner(conn)
	// loop until the connection is closed
	for scanner.Scan() {
		// read the next line of input
		line := scanner.Text()
		// split the line into words
		words := strings.Split(line, " ")
		// process the command
		if len(words) > 0 {
			// check the first word to determine the command
			switch words[0] {
			// disconnect command
			case "disconnect":
				// remove the connection from the server
				srv.rmconn <- conn
				// close the connection
				conn.Close()
				// stop handling the connection
				return

			// any other command
			default:
				// send the command to the server
				srv.commands <- line
			}
		}
	}
	// check for errors
	if err := scanner.Err(); err != nil {
		// handle the error
		fmt.Println(err)
	}
	// remove the connection from the server
	srv.rmconn <- conn
	// close the connection
	conn.Close()
}

// run the server
func (s *server) run() {
	for {
		select {
		// handle new client connections
		case conn := <-s.addconn:
			// store the client connection
			s.connections[conn.RemoteAddr().String()] = conn
			// print a message to the server console
			fmt.Println("added new client connection:", conn.RemoteAddr())

		// handle client disconnections
		case conn := <-s.rmconn:
			// remove the client connection
			delete(s.connections, conn.RemoteAddr().String())
			// print a message to the server console
			fmt.Println("removed client connection:", conn.RemoteAddr())

		// handle client commands
		case cmd := <-s.commands:
			// print the command to the server console
			fmt.Println("received command:", cmd)
			// loop through the connected clients and send the command to each of them
			for _, conn := range s.connections {
				fmt.Fprintln(conn, cmd)
			}
		}
	}
}

func main() {
	// parse the command line arguments
	port := flag.Int("port", 8080, "the port number to listen on")
	flag.Parse()

	// create a new server instance
	srv := newServer()

	// start the server in a goroutine
	go srv.run()

	// listen for incoming client connections on the specified port
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		// handle the error
		fmt.Println(err)
		return
	}

	// print a message to the server console
	fmt.Println("server listening on port", *port)

	// accept incoming client connections in a loop
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle the error
			fmt.Println(err)
			continue
		}

		// add the new client connection to the server
		srv.addconn <- conn

		// start a goroutine to handle the client connection
		go handleConnection(conn, srv)
	}
}
