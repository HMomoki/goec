package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"

	"goec/Models"
)

func Signin(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{
	})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	dbPassword := models.GetUser(email).Password
	formPassword := c.PostForm("password")

	if dbPassword == formPassword {
		c.HTML(http.StatusOK, "login.html", gin.H{
		})
	}else{
		c.Redirect(302,"/signin")
	}

}