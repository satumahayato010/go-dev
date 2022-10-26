package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "top.html", gin.H{})
	})

	r.GET("/page1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page1.html", gin.H{
			"v1": 1,
			"v2": 2,
		})
	})

	r.GET("/page2/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.HTML(http.StatusOK, "page2.html", gin.H{
			"id": id,
		})
	})

	r.GET("/page3", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page3.html", gin.H{})
	})

	r.POST("/page3/submit", func(c *gin.Context) {
		value := c.PostForm("value")
		fmt.Println(value)
		c.Redirect(302, "/page3")
	})

	r.Run(":8080")
}
