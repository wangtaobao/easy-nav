package main

import (
	"easy-nav/models"
	"easy-nav/router"
)

func main() {

	models.NewGormDB()

	r := router.App()

	r.Run(":8081")

}
