package routes

import(
	"github.com/gin-gonic/gin"

	"goec/Controllers"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	
	// auth := router.Group("/auth")
	// {
	router.LoadHTMLGlob("resources/views/*")

	router.GET("/signup",controller.Signup)
	router.POST("/signup",controller.Resister)


	// 	auth.GET("/signup",controller.Signup)
	// 	auth.POST("/signup",controller.Resister)
	// 	auth.GET("/signin",controller.Signin)
	// 	auth.POST("/signin",controller.Login)
	// }

	return router
}