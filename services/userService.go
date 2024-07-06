package services

import (
	"apps/models"
	"apps/repositories"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repositories.GetAllUsers(&users)
	return users, err
}

func CreateUser(user models.User) error {
	return repositories.CreateUser(&user)
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := repositories.GetUserByID(&user, id)
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	return repositories.GetUserByEmail(email)
}

func UpdateUser(user models.User) error {
	return repositories.UpdateUser(&user)
}

func DeleteUser(id string) error {
	var user models.User
	return repositories.DeleteUser(&user, id)
}
