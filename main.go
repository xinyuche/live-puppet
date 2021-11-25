package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// )

// const (
// 	CONN_HOST = "localhost"
// 	CONN_PORT = "3333"
// 	CONN_TYPE = "tcp"
// )

// func main() {
// 	// Listen for incoming connections.
// 	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
// 	if err != nil {
// 		fmt.Println("Error listening:", err.Error())
// 		os.Exit(1)
// 	}
// 	// Close the listener when the application closes.
// 	defer l.Close()
// 	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
// 	for {
// 		// Listen for an incoming connection.
// 		conn, err := l.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting: ", err.Error())
// 			os.Exit(1)
// 		}
// 		// Handle connections in a new goroutine.
// 		go handleRequest(conn)
// 	}
// }

// // Handles incoming requests.
// func handleRequest(conn net.Conn) {
// 	// Make a buffer to hold incoming data.
// 	buf := make([]byte, 1024)
// 	// Read the incoming connection into the buffer.
// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}
// 	// Send a response back to person contacting us.
// 	conn.Write([]byte("Message received."))
// 	// Close the connection when you're done with it.
// 	conn.Close()
// }
import (
	"fmt"
	"log"
	"net/http"
)

// HelloServer responds to requests with the given URL path.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you requested: %s", r.URL.Path)
	log.Printf("Received request for path: %s", r.URL.Path)
}

func main() {
	var addr string = ":8180"
	handler := http.HandlerFunc(HelloServer)
	log.Printf("Starting webserver on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not listen on port %s %v", addr, err)
	}
}
