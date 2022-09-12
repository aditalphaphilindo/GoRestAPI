package model

import (
	"FreshGo/config"

	"golang.org/x/crypto/bcrypt"
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
	var UserFil Users
	var password string = user.Password

	hash, _ := HashPassword(password)

	UserFil.Username = user.Username
	UserFil.Password = hash
	UserFil.Email = user.Email
	UserFil.NamaLengkap = user.NamaLengkap
	UserFil.Alamat = user.Alamat
	UserFil.Tipe = user.Tipe

	if err := config.DB.Create(UserFil).Error; err != nil {
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

func (user *Users) LoginUser(email string) (Users, error) {
	// var UserFil Users
	// if err := config.DB.Where("username = ?", email).First(&UserFil).Error; err != nil {
	// 	return err
	// }
	// return nil
	var userData Users
	result := config.DB.Where("email = ?", email).First(&userData)
	return userData, result.Error
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
