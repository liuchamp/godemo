package models

type User struct {
	User string `json:"user" bson:"user"`
	Name string
}

type UserRegister struct {
	UserName string `json:"userName"`
	Code     string `json:"code"`
}
