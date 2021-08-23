package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"fmt"

	// models "goec/Models"
)

func Signin(c *gin.Context) {
	fmt.Println("New Game")
	c.JSON(http.StatusOK, "New Game")
}

func Login(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	c.JSON(http.StatusOK, name)
	// email := c.PostForm("email")
	// dbPassword := models.GetUser(email).Password
	// formPassword := c.PostForm("password")

	// if dbPassword == formPassword {
	// 	c.HTML(http.StatusOK, "login.html", gin.H{
	// 	})
	// }else{
	// 	c.Redirect(302,"/auth/signin")
	// }

}