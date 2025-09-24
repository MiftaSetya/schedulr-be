package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null" json:"name"`
	Email    string `gorm:"size:150;unique;not null" json:"email"`
	Password string `gorm:"size:250;not null" json:"-"`
}
