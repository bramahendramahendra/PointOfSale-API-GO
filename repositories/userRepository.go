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

func GetUserByID(user *models.User, id uint) error {
	return config.DB.Where("id = ?", id).First(user).Error
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmailTemp(emailTemp string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("email_temp = ?", emailTemp).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return SoftDeleteUser(id)
}

func SoftDeleteUser(id uint) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}
	user.EmailTemp = user.Email
	user.Email = ""
	return config.DB.Save(&user).Error
}

func HardDeleteUser(id uint) error {
	return config.DB.Unscoped().Delete(&models.User{}, id).Error
}
