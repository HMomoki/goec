package routes

import(
	"github.com/gin-gonic/gin"

	"goec/Controllers"
)

func GetRouter() *gin.Engine {
	router := aloowCrossOrigin() 

	auth := router.Group("/auth")
	{
		auth.GET("/signup",controllers.Signup)
		auth.POST("/signup",controllers.Resister)
		auth.GET("/signin",controllers.Signin)
		auth.POST("/signin",controllers.Login)
	}

	return router
}