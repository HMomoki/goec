package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "goec/Models"
	"fmt"
)

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println(email)
	fmt.Println(password)

	models.CreateUser(email,password)

	c.JSON(http.StatusOK,email)
}

// func Signin(c *gin.Context) {
// 	email := c.PostForm("email")
// 	dbPassword := models.GetUser(email).Password
// 	formPassword := c.PostForm("password")

// 	if dbPassword == formPassword {

// 	}else{
	
// 	}
// }