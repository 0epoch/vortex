package service

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"vortex/model"
)

func FindUserByID() {

}

func FindUsers() {

}

func CreateUser (data map[string]interface{}) (*model.User, error) {
	pwd := data["password"].(string)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: data["username"].(string),
		Password:  string(passwordHash),
		LastAt: time.Now(),
	}
	err = model.DB.Create(user).Error
	return user, err
}

