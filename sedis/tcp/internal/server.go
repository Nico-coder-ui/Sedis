package internal

import (
	"fmt"
	"io"
	"net"
)

type Request struct {
	Message string
	Conn    net.Conn
}

func HandleClient(conn net.Conn, ch chan<- Request) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected.")
			} else {
				fmt.Println("Error Read:", err)
			}
			return
		}

		msg := string(buffer[:n])
		ch <- Request{Message: msg, Conn: conn}
	}
}
