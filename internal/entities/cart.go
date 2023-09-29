package entities

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	ID          string        `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	UserId      string        `gorm:"type:varchar(36);not null" json:"user_id"`
	Total       int           `gorm:"type:int;not null" json:"total"`
	CartProduct []CartProduct `json:"products"`
	CreatedAt   time.Time     `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time     `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time    `sql:"index" json:"deletedAt,omitempty"`
}

type CartProduct struct {
	gorm.Model
	ID        string     `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	UserId    string     `gorm:"type:varchar(36);not null" json:"user_id"`
	CartId    string     `gorm:"type:varchar(36);not null" json:"cart_id"`
	ProductId string     `gorm:"type:varchar(36);not null" json:"product_id"`
	Quantity  int        `gorm:"type:int;not null" json:"quantity"`
	Total     int        `gorm:"type:int;not null" json:"total"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
