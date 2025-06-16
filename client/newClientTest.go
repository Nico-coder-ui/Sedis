package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8086")
	if err != nil {
		fmt.Println("Erreur de connexion :", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connecté au serveur. Tape 'quit' pour quitter.")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")

	for {
		if !scanner.Scan() {
			break
		}
		asked := scanner.Text()

		if asked == "quit" {
			fmt.Println("Déconnexion...")
			break
		}

		_, err := conn.Write([]byte(asked))
		if err != nil {
			fmt.Println("Erreur d'envoi :", err)
			break
		}

		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("Déconnecté du serveur.")
					fmt.Println(err)
					os.Exit(0)
				}
				fmt.Print("Serveur :", string(buf[:n]))
				fmt.Print(">>> ")
			}
		}()
	}
}
