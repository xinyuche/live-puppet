package main

import (
	"bytes"
	"log"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST = "0.0.0.0"
	HB_PORT   = "32504"
	CONN_TYPE = "tcp"
)

func main() {
	log.Println("Start Heartbeat Container with update 6")

	// Listen for incoming connections for heartbeat.
	hbl, err := net.Listen(CONN_TYPE, CONN_HOST+":"+HB_PORT)
	if err != nil {
		log.Printf("Error heartbeat listening: %s", err.Error())
		return
	}
	// rangeLower := 10
	// rangeUpper := 30

	// liveDuration := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	liveDuration := 200
	starttime := time.Now()
	log.Printf("Time Duration: %v", liveDuration)

	for {
		// Listen for an incoming connection for Heartbeat.
		hbconn, err := hbl.Accept()
		if err != nil {
			log.Printf("Error accepting Heartbeat: %s", err.Error())
			return
		}
		log.Printf("Listening on Heartbeat Port " + CONN_HOST + ":" + HB_PORT)
		// Handle connections in a new goroutine.
		go handleHeartbeat(hbconn, starttime, liveDuration)
	}

}

// Handles incoming requests.
func handleHeartbeat(conn net.Conn, starttime time.Time, duration int) {
	curtime := time.Now()
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	receivedMsg := string(buf[:])
	// isHeartbeatMsg := strings.Compare(receivedMsg, "heartbeat")
	isHeartbeatMsg := bytes.Compare(buf, []byte("heartbeat"))
	log.Printf("IsHeartbeatMSG: %v", isHeartbeatMsg)
	log.Printf("Received: %s", receivedMsg)
	if err != nil {
		log.Printf("Error reading: %s", err.Error())
	}
	if isHeartbeatMsg == 0 {
		heartbeatResponser(starttime, curtime, duration, conn)
	} else {
		controlResponser(conn)
	}
}

func heartbeatResponser(starttime time.Time, curtime time.Time, liveDuration int, conn net.Conn) {
	log.Println("In heartbeat responser.")
	targetTime := starttime.Add(time.Duration(liveDuration) * time.Second)
	alive := curtime.Before(targetTime)
	log.Printf("Liveness: %v", alive)
	if alive {
		conn.Write([]byte("heartbeat received"))
		log.Println("Response: heartbeat received")
		conn.Close()
	} else {
		conn.Write([]byte("failed"))
		log.Println("Response: failed")
		conn.Close()
	}
}

func controlResponser(conn net.Conn) {
	log.Println("In control responser.")
	log.Println("Response: Control signal received")
	conn.Write([]byte("Control signal received."))
	conn.Close()
	log.Println("sending os.exit(1)")
	os.Exit(1)
}
