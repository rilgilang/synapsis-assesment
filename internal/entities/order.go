package entities

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID            string     `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	TransactionId string     `gorm:"type:varchar(36);not null" json:"transaction_id"`
	ProductId     string     `gorm:"type:varchar(36);not null" json:"category_id"`
	ProductName   string     `gorm:"type:varchar(255);not null" json:"product_name"`
	Quantity      int        `gorm:"type:int;not null" json:"quantity"`
	Total         int        `gorm:"type:int;not null" json:"total"`
	CreatedAt     time.Time  `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
