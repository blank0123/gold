package models

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string
	Password string
}

// 判断用户名是否存在
func IsExistUsername(username string) (bool, error) {
	var count int
	err := DB.Model(User{}).Where("username = ?", username).Count(&count).Error
	if count > 0 {
		return true, err
	}
	return false, err
}

func IsExistID(id string) (bool, error) {
	var count int
	err := DB.Model(User{}).Where("id = ?", id).Count(&count).Error
	if count > 0 {
		return true, err
	}
	return false, err
}

func SaveUser(username, password string) error {
	return DB.Save(&User{
		Username: username,
		Password: password,
	}).Error
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.Model(&User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByID(id string) (*User, error) {
	var user User
	err := DB.Model(&User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func EditPasswordByUsername(username, password string) error {
	return DB.Model(&User{}).Where("username = ?", username).Update("password", password).Error
}

func DeleteUserByID(id string) error {
	return DB.Delete(&User{}, id).Error
}

func GetUsers(page, size int)([]User, error){
	var users []User
	err := DB.Model(&User{}).Limit(size).Offset(page).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}