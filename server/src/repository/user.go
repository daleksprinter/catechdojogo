package repository

import (
	"github.com/daleksprinter/catechdojo/db"
	"github.com/daleksprinter/catechdojo/model"
)

func CreateUser(name string, token string) error {
	user := model.User{
		Name:  name,
		Token: token,
	}

	err := db.DB.Create(&user).Error
	return err
}

func GetUser(token string) (model.User, error) {
	user := model.User{}
	err := db.DB.First(&user, "token = ?", token).Error

	return user, err
}

func UpdateUser(token string, name string) error {
	userOld := model.User{}

	userNew := userOld
	userNew.Name = name

	err := db.DB.First(&userOld, "token = ?", token).Error

	if err != nil {
		return err
	}

	err = db.DB.Model(&userOld).Update(&userNew).Error
	if err != nil {
		return err
	}

	return nil
}
