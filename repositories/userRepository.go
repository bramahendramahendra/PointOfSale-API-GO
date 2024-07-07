package repositories

import (
	"apps/config"
	"apps/models"
	"strconv"
)

// Get All data user
func GetAllUsers(users *[]models.User) error {
	return config.DB.Find(users).Error
}

// Creta data user
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// Get data user berdasarkan ID
func GetUserByID(user *models.User, id uint) error {
	return config.DB.Where("id = ?", id).First(user).Error
}

// Get data user berdasarkan email
func GetUserByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).First(&user).Error
}

// Get data user berdasarkan email temp
func GetUserByEmailTemp(user *models.User, emailTemp string) error {
	return config.DB.Where("email_temp = ?", emailTemp).First(&user).Error
}

// Update user
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// Soft Delete User
func SoftDeleteUser(id uint) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}
	user.EmailTemp = user.Email
	user.Email = "deleted_" + strconv.Itoa(int(user.ID))

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return config.DB.Delete(&user).Error
}

// Hard Delete User
func HardDeleteUser(id uint) error {
	return config.DB.Unscoped().Delete(&models.User{}, id).Error
}
