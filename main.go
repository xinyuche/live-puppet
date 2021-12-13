package main

import (
	"log"
	"net"
	"os"
)

const (
	CONN_HOST    = "0.0.0.0"
	CONTROL_PORT = "32505"
	CONN_TYPE    = "tcp"
)

func main() {
	log.Println("Start Control Container")
	// Listen for incoming connections for operation.
	ctllistener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONTROL_PORT)
	if err != nil {
		log.Printf("Error operation listening: %s", err.Error())
		return
	}
	// Close the operation listener when the application closes.
	defer ctllistener.Close()

	go func() {
		for {
			// Listen for an incoming connection for Heartbeat.
			ctlconn, err := ctllistener.Accept()
			if err != nil {
				log.Printf("Error accepting control signal: %s", err.Error())
				return
			}
			log.Printf("Listening on control port " + CONN_HOST + ":" + CONTROL_PORT)
			// Handle connections in a new goroutine.
			go handleControlSignal(ctlconn)
		}
	}()
}

// Handles incoming requests.
func handleControlSignal(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	log.Printf("Received Control Signal: %s", string(buf[:]))
	if err != nil {
		log.Printf("Error reading: %s", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Control Signal received."))
	// Close the connection when you're done with it.
	conn.Close()
	os.Exit(1)
}
