package main

import (
	"net/http"
	c "shorturl/controller"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":     0,
			"message":    "pong",
			"serverTime": time.Now(),
		})
	})
	//v1 api
	{
		r := r.Group("/v1")
		r.GET("/gen", c.Gen)
		r.GET("/get", c.Get)
		r.GET("/list", c.List)
	}
	r.Run(":9999")
}
