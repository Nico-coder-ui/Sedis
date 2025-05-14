package main

import (
	"fmt"
	"net"
	"strings"

	"mini-redis/internal"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error Listen:", err)
		return
	}
	defer listener.Close()

	ch := make(chan internal.Request, 5)
	s := internal.NewStore()
	s.StartExpirationLoop()
	s.Load("store.json")

	go func() {
		for req := range ch {
			fmt.Printf("Received on server : %s\n", strings.TrimSpace(req.Message))
			internal.ParseMsg(req.Message, &s, req.Conn)
		}
	}()

	fmt.Println("Mini-Redis Server running on localhost:8000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accept:", err)
			continue
		}

		go internal.HandleClient(conn, ch)
	}
}
