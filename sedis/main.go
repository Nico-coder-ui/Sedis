package main

import (
	"fmt"
	routes "sedis/http/routes"
	"sedis/stokage"
	"sedis/tcp"
)

func main() {
	store := stokage.NewStore()
	store.Load("json/store.json")
	store.StartExpirationLoop()

	errChan := make(chan error, 2)

	router := routes.SetupRouter(&store)

	go func() {
		errChan <- router.Run("0.0.0.0:8085")
	}()

	go tcp.Start(":8086", &store, errChan)

	if err := <-errChan; err != nil {
		fmt.Printf("Server error: %v", err)
	}
}
