package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8180"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Printf("Error listening: %s", err.Error())
		return
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Printf("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting: %s", err.Error())
			return
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	log.Printf("received: %s", string(buf[:]))
	if err != nil {
		log.Printf("Error reading: %s", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
