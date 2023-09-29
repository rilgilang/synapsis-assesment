package migrations

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/entities"
)

var models = []interface{}{
	&entities.User{},
	&entities.Product{},
	&entities.Cart{},
	&entities.CartProduct{},
	&entities.Transaction{},
	&entities.ProductCategory{},
	&entities.Order{},
}

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
}
