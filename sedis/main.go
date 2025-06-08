package main

import (
	routes "sedis/http/routes"
	"sedis/stokage"
	"sedis/tcp"
)

func main() {
	store := stokage.NewStore()
	store.StartExpirationLoop()
	store.Load("store.json")

	router := routes.SetupRouter(&store)
	router.Run(":8085")

	tcp.Start(":8086", &store)
}
