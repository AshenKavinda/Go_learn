package migration

import (
	"log"

	"github.com/ashenkavinda/go_social_app/internel/models"
	"gorm.io/gorm"
)

func SqlMigration(gorm *gorm.DB) {
	err := gorm.AutoMigrate(&models.User{}, &models.Post{}, &models.Follow{})
	if err != nil {
		log.Fatalf("auto migrate error: %v", err)
		return
	}

	log.Println("Auto migration completed")
}
