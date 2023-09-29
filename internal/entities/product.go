package entities

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID              string          `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	ProductName     string          `gorm:"type:varchar(255);not null" json:"product_name"`
	CategoryId      string          `gorm:"type:varchar(36);not null" json:"category_id"`
	ProductCategory ProductCategory `gorm:"foreignKey:CategoryId"`
	Price           int             `gorm:"type:int;not null" json:"price"`
	CreatedAt       time.Time       `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time       `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt       *time.Time      `sql:"index" json:"deletedAt,omitempty"`
}

type ProductCategory struct {
	gorm.Model
	ID           string     `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	CategoryName string     `gorm:"type:varchar(255)" json:"category"`
	CreatedAt    time.Time  `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
