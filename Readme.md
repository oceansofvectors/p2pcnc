# P2P Command and Control Server

This is a simple peer-to-peer command and control server written in Go. The server allows clients to connect and send commands that are broadcasted to all other connected clients. The server maintains a list of active client connections and uses channels to communicate with the client goroutines.
Features

* Client connections are stored in a map keyed by IP address
* Commands are received and broadcasted using channels
* Client connections can be added and removed using channels
* The server listens for incoming client connections on port 8080
* The server can handle multiple client connections concurrently

## Usage

To use the server, you first need to have Go installed on your system. Once you have Go installed, you can download and build the server using the following commands:

`go build server.go`

`./server`

This will start the server and listen for incoming client connections on port 8080. To connect to the server, you can use a telnet client or any other tool that can create a TCP connection. For example, you can use the telnet command to connect to the server:

`telnet localhost 8080`

Once you are connected to the server, you can send commands by typing them into the terminal and pressing the enter key. The server will receive the command and broadcast it to all connected clients.

This project is still in development and not complete. It is a work in progress and currently only implements a basic set of features for a peer-to-peer command and control server. There are many improvements and additional features that could be added to the server to make it more robust and versatile. For example, the server could be extended to support the following features:

* Add support for authentication and encryption
* Add additional command options and command-specific responses
* Integrate with other tools and systems to provide a more comprehensive solution
* Implement additional features and improvements as needed