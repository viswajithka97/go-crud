package main

import (
	"crud/initializers"
	"crud/models"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
