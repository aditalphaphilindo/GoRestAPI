package main

import (
	"FreshGo/config"
	"FreshGo/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// konek ke database
	config.ConnectDB()
	// route object
	route := echo.New()
	// route list
	route.GET("/", HomePage)
	route.GET("/user/get", GetUserData)
	route.POST("user/create", CreateUserData)
	route.PUT("user/update/:email", UpdateUserData)
	route.DELETE("user/delete/:email", DeleteUserData)

	// auth route
	route.POST("/auth/login/process", AuthLoginProcess)

	route.Start(":9000")
}

func HomePage(c echo.Context) error {
	data := "<pre>Homepage</pre>"
	return c.HTML(http.StatusOK, data)
}

func AuthLoginProcess(c echo.Context) error {
	user := new(model.Users)
	c.Bind(user)
	response := new(Response)

	return c.JSON(http.StatusOK, response)
}

func UpdateUserData(c echo.Context) error {
	user := new(model.Users)
	c.Bind(user)
	response := new(Response)
	if user.UpdateUser(c.Param("email")) != nil {
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
	user, _ := model.GetOneByEmail(c.Param("email"))
	response := new(Response)

	if user.DeleteUser() != nil {
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
	}
	response := new(Response)
	if user.CreateUser() != nil {
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

	users, err := model.GetAll(c.QueryParam("keywords"))
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
