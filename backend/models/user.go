package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"` // We won't expose the password in JSON
	Token     string    `gorm:"type:varchar(255)" json:"token"`
	CartID    uint      `json:"cart_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}