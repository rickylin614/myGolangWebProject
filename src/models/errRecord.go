package models

type ErrRecord struct {
	Id        uint `gorm:"primary_key"`
	Name      string
	UserId    uint
	ErrMsg    string
	UserAgent string
	Ip        string
}
