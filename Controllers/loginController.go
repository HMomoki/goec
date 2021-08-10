package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"

	"goec/Models"
	"fmt"
)

func Signin(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{
	})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	dbPassword := models.GetUser(email).Password
	formPassword := c.PostForm("password")

	fmt.Println("db=")
	fmt.Println(dbPassword)
	fmt.Println("form=")
	fmt.Println(formPassword)

	if dbPassword == formPassword {
		c.HTML(http.StatusOK, "login.html", gin.H{
		})
	}else{
		c.Redirect(302,"/auth/signin")
	}

}