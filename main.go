package main

import (
	"FreshGo/config"
	"FreshGo/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	// konek ke database
	config.ConnectDB()
	// route object
	route := echo.New()
	// route list
	route.GET("/user/get", GetUserData)
	route.POST("user/create", CreateUserData)
	route.PUT("user/update/:email", UpdateUserData)
	route.DELETE("user/delete/:email", DeleteUserData)
	route.Start(":9000")
}

func UpdateUserData(c echo.Context) error {
	user := new(model.Users)
	c.Bind(user)
	response := new(Response)
	if user.UpdateUser(c.Param("email")) != nil { // method update user
		response.ErrorCode = 10
		response.Message = "Gagal update data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses update data user"
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteUserData(c echo.Context) error {
	user, _ := model.GetOneByEmail(c.Param("email")) // method get by email
	response := new(Response)

	if user.DeleteUser() != nil { // method update user
		response.ErrorCode = 10
		response.Message = "Gagal menghapus data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses menghapus data user"
	}
	return c.JSON(http.StatusOK, response)
}

func CreateUserData(c echo.Context) error {
	user := new(model.Users)
	c.Bind(user)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		// fmt.Println("Request dari json")
	} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
		file, err := c.FormFile("ktp")
		if err != nil {
			fmt.Println("File tidak terupload")
		} else {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			dst, err := os.Create(file.Filename)
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			user.Ktp = file.Filename
			fmt.Println("Ada file, akan disimpan")
		}
	}
	response := new(Response)
	if user.CreateUser() != nil { // method create user
		response.ErrorCode = 10
		response.Message = "Gagal create data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create data user"
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)
}

func GetUserData(c echo.Context) error {
	response := new(Response)

	users, err := model.GetAll(c.QueryParam("keywords")) // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal load data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses load data user"
		response.Data = users
	}
	fmt.Println(response)
	return c.JSON(http.StatusOK, response)
}

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}
