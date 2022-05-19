package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		string `gorm:"column:name"`
	Age  		uint   `gorm:"column:age"`
}