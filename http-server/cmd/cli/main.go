package main

import (
	"fmt"
	poker "github.com/keigodasu/learn-go-with-tests/http-server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem getting store from file %v", err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
