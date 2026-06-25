package main

import (
	"log"
	"os"

	"github.com/ashenkavinda/go_social_app/internel/store"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("port")

	store := store.NewPostgresStorage(nil)

	cfg := config{
		addr: ":" + port,
	}

	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
