package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("port")

	cfg := config{
		addr: ":" + port,
	}

	app := application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
