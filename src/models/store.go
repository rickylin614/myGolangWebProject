package models

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	// Id        int64
	Name        string
	Phone_no    int
	Region      string
	Create_user string
	Update_user string
	//gorm.Model 其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	gorm.Model
}

type StoreResponse struct {
	ID          uint
	Name        string
	PhoneNo     string
	Region      string
	Create_user string
	CreatedAt   string
}
