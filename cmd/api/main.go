package main

import (
	"log"

	"github.com/ashenkavinda/go_social_app/internel/config"
	"github.com/ashenkavinda/go_social_app/internel/db"
	"github.com/ashenkavinda/go_social_app/internel/db/migration"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()

	gorm, sqlDB := db.InitializeConnection(cfg.Server.DSN)
	defer sqlDB.Close()

	cfg.Server.DB = gorm

	migration.SqlMigration(gorm)

	app := Application{
		Config: cfg,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
