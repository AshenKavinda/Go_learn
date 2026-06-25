package db

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeConnection(dsn string) (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get db error: %v", err)
		return nil, nil
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("ping db error: %v", err)
		return nil, nil
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection successful")

	return db, sqlDB
}
