package controller

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"app": "GO RestFull API V.1",
	})
}
