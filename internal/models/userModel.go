package models

import (
	"mangosteen/internal/database"
)

type User struct {
	ID         int     `gorm:"primary_key" json:"id"`
	Name       string  `gorm:"type:varchar(20);not null" json:"username"`
	Email      *string `gorm:"type:varchar(20);unique_index" json:"email"`
	Age        *uint8  `gorm:"type:int" json:"age"`
	mobile     string  `gorm:"type:varchar(11);unique_index" json:"mobile"`
	Address    string  `gorm:"type:varchar(150);index:addr" json:"address"`
	CreateTime int64   `gorm:"type:int" json:"create_time"`
	UpdateTime int64   `gorm:"type:int" json:"update_time"`
}

func CreateUserTest() {
	database.GormDB.AutoMigrate(&User{})
}
