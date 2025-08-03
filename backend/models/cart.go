package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	Items     []Item    `gorm:"many2many:cart_items;" json:"items"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}