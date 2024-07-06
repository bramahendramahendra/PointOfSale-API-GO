package repositories

import (
	"apps/config"
	"apps/models"
)

func GetAllUsers(users *[]models.User) error {
	return config.DB.Find(users).Error
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByID(user *models.User, id string) error {
	return config.DB.Where("id = ?", id).First(user).Error
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(user *models.User, id string) error {
	return config.DB.Where("id = ?", id).Delete(user).Error
}
