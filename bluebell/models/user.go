package models

import "time"

// User --> 数据库表
type User struct {
	ID        int64 `gorm:"primary_key"`
	UserId    int64 `gorm:"not null"`
	Username  string
	Password  string
	Email     string
	Gender    int8 `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
