package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"

	"goec/Models"
)

func Signup(c *gin.Context){
	title := "title"
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"title" : title,
	})
}

func Resister(c *gin.Context){
	email := c.PostForm("email")
	password := c.PostForm("password")

	models.CreateUser(email,password)

	c.HTML(http.StatusOK, "resister.html", gin.H{
		"email" : email,
		"password" : password,
	})
}