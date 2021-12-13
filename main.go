package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "0.0.0.0"
	HB_PORT   = "32504"
	CONN_TYPE = "tcp"
)

func main() {
	log.Println("Start Heartbeat Container")
	// Listen for incoming connections for heartbeat.
	hbl, err := net.Listen(CONN_TYPE, CONN_HOST+":"+HB_PORT)
	if err != nil {
		log.Printf("Error heartbeat listening: %s", err.Error())
		return
	}

	// Close the heartbeat listener when the application closes.
	defer hbl.Close()

	for {
		// Listen for an incoming connection for Heartbeat.
		hbconn, err := hbl.Accept()
		if err != nil {
			log.Printf("Error accepting Heartbeat: %s", err.Error())
			return
		}
		log.Printf("Listening on Heartbeat Port " + CONN_HOST + ":" + HB_PORT)
		// Handle connections in a new goroutine.
		go handleHeartbeat(hbconn)
	}

}

// Handles incoming requests.
func handleHeartbeat(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	log.Printf("Received Heartbeat: %s", string(buf[:]))
	if err != nil {
		log.Printf("Error reading: %s", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Heartbeat received."))
	// Close the connection when you're done with it.
	conn.Close()
}
