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
	// Check if the email exists in the temp_email field
	existingUser, err := repositories.GetUserByEmailTemp(user.Email)
	if err == nil && existingUser.ID != 0 {
		// If the existing user was soft deleted, update it with new data
		existingUser.Email = user.Email
		existingUser.EmailTemp = ""
		existingUser.Name = user.Name
		return repositories.UpdateUser(&existingUser)
	}
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

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}

func HardDeleteUser(id uint) error {
	return repositories.HardDeleteUser(id)
}
