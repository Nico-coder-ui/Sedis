package tcp

import (
	"fmt"
	"net"
	"strings"

	"sedis/stokage"
	"sedis/tcp/internal"
)

func Start(port string, s *stokage.Store) {
	conn, err := net.Listen("tcp", "localhost"+port)
	if err != nil {
		fmt.Println("Error Listen:", err)
		return
	}
	defer conn.Close()

	ch := make(chan internal.Request, 5)

	go func() {
		for req := range ch {
			fmt.Printf("Received on server : %s\n", strings.TrimSpace(req.Message))
			internal.ParseMsg(req.Message, s, req.Conn)
		}
	}()

	fmt.Println("Mini-Redis Server running on localhost:8000...")

	for {
		conn, err := conn.Accept()
		if err != nil {
			fmt.Println("Error Accept:", err)
			continue
		}

		go internal.HandleClient(conn, ch)
	}
}
