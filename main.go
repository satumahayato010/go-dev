package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/user", QiitaGet)
	r.Run(":8080")
}