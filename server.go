package main

import(
  "github.com/gin-gonic/gin"
  "fmt"
)

func main() {
	api := gin.Default()
	api.Use(DummyMiddleware())

	v1 := router.Group("/v1")
	{
		v1.GET("/customer", getting)
		v1.POST("/customer", posting)
		v1.DELETE("/customer", deleting)
	}

	api.Run()
}