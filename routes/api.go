package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func aloowCrossOrigin() *gin.Engine {
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

	return router
}