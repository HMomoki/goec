package routes

import(
	"github.com/gin-gonic/gin"

	"goec/Controllers"
)

func GetRouter() *gin.Engine {
	router := aloowCrossOrigin() 

	auth := router.Group("/auth")
	{
		auth.POST("/signup",controllers.Signup)
	}

	return router
}