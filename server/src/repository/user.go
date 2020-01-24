package repository

import (
	"github.com/daleksprinter/catechdojo/model"
	"github.com/daleksprinter/catechdojo/db"
)

func CreateUser(name string, token string) model.User {
	user := model.User{
		Name: name,
		Token: token,
	}

	db.DB.Create(&user)
	return user
}

func GetUser(token string) model.User {
	user := model.User{}
	user.Token = token
	db.DB.First(&user)

	return user
}
