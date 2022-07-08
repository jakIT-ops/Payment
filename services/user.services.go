package services

import (
	"payment_full/db"
	"payment_full/models"
)

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
}

func (u *User) CreateUser(arg CreateUserParams) (models.User, error) {
	var newUser models.User

	newUser.UserName = arg.Username
	newUser.Password = arg.Password
	newUser.Email = arg.Email

	db.Database.Db.Create(&newUser)

	return newUser, nil
}

func (u *User) GetUser(username string) (models.User, error) {
	var user models.User

	err := db.Database.Db.Find(&user, "user_name = ?", username)

	if err.Error != nil {
		return user, nil
	}
	return user, err.Error
}
