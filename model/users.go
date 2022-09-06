package model

import (
	"FreshGo/config"
	"fmt"
)

type Users struct {
	Id          int
	Username    string `json:"username" form:"username" gorm:"primaryKey"`
	Password    string `json:"password" form:"password"`
	Email       string `json:"email" form:"email"`
	NamaLengkap string `json:"nama" form:"nama"`
	Alamat      string `json:"alamat" form:"alamat"`
	Tipe        string `json:"tipe" form:"tipe"`
}

func (user *Users) CreateUser() error {
	fmt.Println(user.Password)
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) UpdateUser(email string) error {
	if err := config.DB.Model(&Users{}).Where("email = ?", email).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteUser() error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneByEmail(email string) (Users, error) {
	var user Users
	result := config.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func GetAll(keywords string) ([]Users, error) {
	var users []Users
	result := config.DB.Find(&users)
	return users, result.Error
}
