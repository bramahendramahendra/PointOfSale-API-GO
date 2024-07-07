package services

import (
	"apps/models"
	"apps/repositories"
)

// Get All data user
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repositories.GetAllUsers(&users)
	return users, err
}

// Creta data user
func CreateUser(user models.User) error {
	return repositories.CreateUser(&user)
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := repositories.GetUserByID(&user, id)
	return user, err
}

// Get data user berdasarkan email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := repositories.GetUserByEmail(&user, email)
	return user, err
}

// Update user
func UpdateUser(user models.User) error {
	return repositories.UpdateUser(&user)
}

// Soft Delete User
func SoftDeleteUser(id uint) error {
	return repositories.SoftDeleteUser(id)
}

// Hard Delete User
func HardDeleteUser(id uint) error {
	return repositories.HardDeleteUser(id)
}
