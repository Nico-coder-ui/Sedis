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

func HandleClient(conn net.Conn, ch chan<- Request, errChan chan<- error) {
	fmt.Println("Client détecté")
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		fmt.Println("Read")
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected.")
			} else {
				errChan <- fmt.Errorf("error read:%s", err)
				fmt.Println("Erreur:", err)
			}
			return
		}
		fmt.Println("Message enregistré")
		msg := string(buffer[:n])
		ch <- Request{Message: msg, Conn: conn}
	}
}
