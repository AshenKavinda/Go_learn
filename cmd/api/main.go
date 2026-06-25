package main

import (
	"log"

	"github.com/ashenkavinda/go_social_app/internel/app"
	"github.com/ashenkavinda/go_social_app/internel/config"
	"github.com/ashenkavinda/go_social_app/internel/db"
	"github.com/ashenkavinda/go_social_app/internel/db/migration"
	"github.com/ashenkavinda/go_social_app/internel/store"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()

	gorm, sqlDB := db.InitializeConnection(cfg.DSN)
	defer sqlDB.Close()

	migration.SqlMigration(gorm)

	store := store.NewPostgresStorage(sqlDB)

	app := app.Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
