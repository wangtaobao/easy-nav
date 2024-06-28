package router

import (
	"easy-nav/middleware"
	"easy-nav/service"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.POST("/api/add", service.AddUrl)

	r.GET("/api", service.GetIndexList)

	r.DELETE("/api/edit/:id", service.DelIndexById)

	r.PUT("/api/edit", service.UpdateIndex)

	return r
}
