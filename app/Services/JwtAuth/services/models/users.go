package models

type User struct {
	Id     string `json:"id" form:"-"`
	Name string `json:"name" form:"username"`
	Password string `json:"password" form:"password"`
}
