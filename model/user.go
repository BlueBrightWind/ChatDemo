package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	PassWord string
	Phone    string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email    string `valid:"email"`
	Avatar   string //头像
	Salt     string
}
