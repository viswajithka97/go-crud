package main

import (
	"crud/controllers"
	"crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvFile()
	initializers.ConnectDB()
}

func main() {

	r := gin.Default()
	r.POST("/createPosts", controllers.CreatePosts)
	r.GET("/getAllPosts", controllers.FetchAllPosts)
	r.GET("/postById/:id", controllers.PostById)
	r.PUT("/updatePost/:id", controllers.PostUpdate)
	r.DELETE("/deletePost/:id", controllers.PostDelete)

	r.Run()
}
