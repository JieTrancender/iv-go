package models

import "github.com/jinzhu/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

// AddUser 新增用户
func AddUser(user *User) {
	DB.Create(&user)
	return
}

// UserDetailByName 使用用户名查找用户
func UserDetailByName(name string) (user User) {
	DB.Where("name = ?", name).First(&user)
	return
}

// UserDetailByEmail 使用邮件查找用户
func UserDetailByEmail(email string) (user User) {
	DB.Where("email = ?", email).First(&user)
	return
}

// UserDetail 使用用户id查找用户
func UserDetail(id uint) (user User) {
	DB.Where("id = ?", id).First(&user)
	return
}

// GetUserTotal 获取总用户数
func GetUserTotal() (int, error) {
	var count int
	if err := DB.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetUsers 获取指定页面用户
func GetUsers(page int, size int) (users []User) {
	DB.Offset((page - 1) * size).Limit(size).Find(&users)
	return
}
