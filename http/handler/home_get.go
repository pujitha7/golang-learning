package handler

import "github.com/gin-gonic/gin"

func HomeGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Yo",
	})
}
