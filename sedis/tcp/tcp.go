package tcp

import (
	"fmt"
	"net"

	"sedis/stokage"
	"sedis/tcp/internal"
)

func Start(port string, s *stokage.Store, errChan chan<- error) {
	conn, err := net.Listen("tcp", "0.0.0.0"+port)
	if err != nil {
		errChan <- fmt.Errorf("error listening on %s: %w", port, err)
		return
	}
	defer conn.Close()

	ch := make(chan internal.Request, 5)

	fmt.Println("Lancement goroutine")
	go func() {
		for req := range ch {
			internal.ParseMsg(req.Message, s, req.Conn)
		}
	}()

	fmt.Println("Mini-Redis Server running on localhost" + port + "...")

	fmt.Println("Lancement boucle")
	for {
		conn, err := conn.Accept()
		if err != nil {
			errChan <- fmt.Errorf("error accept:%s", err)
			fmt.Println("Erreur:", err)
			continue
		}
		fmt.Println("Pas d'erreur")
		go internal.HandleClient(conn, ch, errChan)
	}
}
