package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening on port 8080")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	// Schickt eine Begrüßungsnachricht an den Client
	conn.Write([]byte("Hello from server\n"))

	// Erstellt einen Buffer-Reader, um die Eingaben des Clients zu lesen
	buffer := bufio.NewReader(conn)

	// Endlos-Schleife, um Eingaben des Clients zu lesen
	for {
		// Liest die Eingabe des Clients bis zum Newline-Charakter
		clientInput, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			break
		}

		// Ausgabe der Eingabe des Clients auf der Konsole
		fmt.Printf("Client says: %s", clientInput)

		// Schreibt die Eingabe zurück zum Client
		conn.Write([]byte("You said: " + clientInput))
	}
}
