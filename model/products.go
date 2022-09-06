package model

import "FreshGo/config"

type Products struct {
	Id           int
	Kd_product   string `json:"kd_product" form:"kd_product"`
	Nama_product string `json:"nama_produk" form:"nama_produk"`
	Deks         string `json:"deks" form:"deks"`
	Harga        int    `json:"harga" form:"harga"`
}

func GetAllDataProduct(keywords string) ([]Users, error) {
	var users []Users
	result := config.DB.Find(&users)

	return users, result.Error
}
