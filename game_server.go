package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func StartGameServer() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	buf := make([]byte, 1024) // Create empty

	reqLen, err := conn.Read(buf) // Read the incoming connection into the buffer.
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	// Print message received
	fmt.Printf("Message contents: %q\n", string(buf[:reqLen]))

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))

	// Close the connection when you're done with it.
	conn.Close()
}

func StartClient(msg string) {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(msg)); err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Read error: %s", err)
			}
			break
		}
		fmt.Printf("Message contents: %q\n", string(buf[:reqLen]))
	}

}
