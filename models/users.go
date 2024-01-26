package models

import (
	"materialsSort/global"
	"materialsSort/models/res"
)

type Users struct {
	UserId   int    `gorm:"primary_key;AUTO_INCREMENT" json:"user_id"`
	UserName string `gorm:"type:varchar(20);not null" json:"user_name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Email    string `gorm:"type:varchar(20);not null" json:"email"`
}

func (Users) TableName() string {
	return "users"
}

// AddUser 添加用户
func AddUser(userName string, password string, email string) (int, error) {
	user := Users{UserName: userName, Password: password, Email: email}
	if err := global.DB.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.UserId, nil
}

// GetUserById 根据用户id获取用户
func GetUserById(userId int) (*res.User, error) {
	user := Users{}
	if err := global.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	userResponse := res.User{UserId: user.UserId, UserName: user.UserName, Email: user.Email}
	return &userResponse, nil
}

// GetUserByName 根据用户名获取用户
func GetUserByName(userName string) (*res.User, error) {
	user := Users{}
	if err := global.DB.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	userResponse := res.User{UserId: user.UserId, UserName: user.UserName, Email: user.Email}
	return &userResponse, nil
}

// UpdateUser 更新用户
func UpdateUser(userId int, userName string, password string, email string) error {
	user := Users{UserId: userId, UserName: userName, Password: password, Email: email}
	if err := global.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func DeleteUser(userId int) error {
	user := Users{}
	if err := global.DB.Where("user_id = ?", userId).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetAllUsers 获取所有的用户，不分页
func GetAllUsers() (int64, []*res.User, error) {
	var count int64
	var users []*res.User
	if err := global.DB.Model(&Users{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}
	if err := global.DB.Find(&users).Error; err != nil {
		return 0, nil, err
	}
	return count, users, nil
}

// GetAllUsersWithPage 获取所有的用户，分页
func GetAllUsersWithPage(page int, pageSize int) (int64, []*res.User, error) {
	var count int64
	var users []*res.User
	if err := global.DB.Model(&Users{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}
	if err := global.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return 0, nil, err
	}
	return count, users, nil
}
