package main

import (
	"goec/routes"
)

func main(){
	router := routes.GetRouter()

	router.Run(":8000")
}