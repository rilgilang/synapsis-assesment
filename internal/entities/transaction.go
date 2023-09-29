package entities

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	ID        string     `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	UserId    string     `gorm:"type:varchar(36);not null" json:"user_id"`
	Total     int        `gorm:"type:int;not null" json:"total"`
	Status    string     `gorm:"type:enum('waiting', 'failed', 'completed');default:waiting" json:"status"`
	Order     []Order    `json:"orders"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
