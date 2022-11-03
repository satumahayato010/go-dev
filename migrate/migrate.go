package main

import (
	"go-dev/initializers"
	"go-dev/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
