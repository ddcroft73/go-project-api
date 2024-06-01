package model

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey"`
	Username    string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	IsSuperUser bool      `gorm:"default:false"`    // Super User is only created by abother super user.
	Email       string    `gorm:"unique;not null"`
	Phone       string    `gorm:"unique;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
