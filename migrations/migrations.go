package migrations

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/entities"
)

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entities.User{})
}
