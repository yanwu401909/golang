package controller

import (
	"log"
	"net/http"
	"shorturl/service"
	"time"

	"github.com/gin-gonic/gin"
)

var serviceImpl *service.MapperService

func init() {
	log.Println("controller init ...")
	serviceImpl = service.BuildMapService()
}

func Gen(c *gin.Context) {
	url := c.DefaultQuery("url", "")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":     http.StatusBadRequest,
			"message":    "Param url should not be blank!",
			"serverTime": time.Now(),
		})
	} else {
		shortUrl, _ := serviceImpl.Generate(url)
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"message":    "success",
			"data":       shortUrl,
			"serverTime": time.Now(),
		})
	}
}

func Get(c *gin.Context) {
	short := c.DefaultQuery("shortUrl", "")
	if short == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":     http.StatusBadRequest,
			"message":    "Param short should not be blank!",
			"serverTime": time.Now(),
		})
	} else {
		url, err := serviceImpl.GetShortUrl(short)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":     http.StatusBadRequest,
				"message":    err.Error(),
				"serverTime": time.Now(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":     http.StatusOK,
				"message":    "success",
				"data":       url,
				"serverTime": time.Now(),
			})
		}
	}
}

func List(c *gin.Context) {
	list, _ := serviceImpl.List()
	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"message":    "success",
		"data":       list,
		"serverTime": time.Now(),
	})
}
