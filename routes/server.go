package routes

import(
	"github.com/gin-gonic/gin"

	"goec/Controllers"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	
	auth := router.Group("/auth")
	{
		router.LoadHTMLGlob("resources/views/auth/*")

		auth.GET("/signup",controllers.Signup)
		auth.POST("/signup",controllers.Resister)
		auth.GET("/signin",controllers.Signin)
		auth.POST("/signin",controllers.Login)
	}

	return router
}