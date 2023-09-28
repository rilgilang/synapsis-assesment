package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        int        `gorm:"primary_key" json:"id"`
	Username  string     `gorm:"type:varchar(255)" json:"username"`
	Password  string     `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
