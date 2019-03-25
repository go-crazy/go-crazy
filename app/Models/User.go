package Model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"username"`
	Password string `json:"password" form:"password"`
}

func (User) TableName() string {
	return "users"
}
