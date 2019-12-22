package main

import (
	"github.com/pwinning1991/golang_with_tests/time"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
