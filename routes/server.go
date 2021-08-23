package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"goec/Controllers"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	auth := router.Group("/auth")
	{
		auth.GET("/signup",controllers.Signup)
		auth.POST("/signup",controllers.Resister)
		auth.GET("/signin",controllers.Signin)
		auth.POST("/signin",controllers.Login)
	}

	return router
}