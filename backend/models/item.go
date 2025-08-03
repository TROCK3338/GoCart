package models

import "time"

type Item struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}