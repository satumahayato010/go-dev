package main

import (
	"go-dev/initializers"
	"go-dev/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}
