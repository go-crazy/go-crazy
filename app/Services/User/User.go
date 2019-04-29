package UserService

import (
	"go-crazy/app"
	"go-crazy/app/Models"
)

func GetUserByName(name string) *Model.User {
	var user Model.User
	App.DB().Where("name = ?", name).First(&user)
	return &user
}

func GetUserByID(id uint) *Model.User {
	var user Model.User
	App.DB().Where("id = ?", id).First(&user)
	if user.ID == id {
		return &user
	}
	return nil
}
