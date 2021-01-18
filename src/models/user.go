package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	// Id        int64
	Name      string
	Pwd       string
	SessionId string
	LoginTime *time.Time `gorm:"type:time"`
	gorm.Model
}

type UserResponse struct {
	ID        uint
	Name      string
	LoginTime string
	CreatedAt string
}
