package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        string     `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	Username  string     `gorm:"type:varchar(255);not null" json:"username"`
	Password  string     `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt;not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt;not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
