package main

import (
	"go-learning/http/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.HomeGet)
	r.Run()
}
