package main

import (
	"github.com/adeben33/golangMiniProject/Todo/controllers"
	"github.com/adeben33/golangMiniProject/Todo/initializers"
	"github.com/adeben33/golangMiniProject/Todo/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	authorized := r.Group("/")
	authorized.Use(middleware.Userauthoraization)
	{
		//	//Get all todo
		authorized.GET("/getTodo", controllers.GetTodo)
		//
		//	//Get a todo by id
		authorized.GET("/getTodo/:id", controllers.GetTodoById)
		//
		//	//Insert a todo
		//	authorized.POST()
		//
		//	//Delete a todo by id
		//	authorized.DELETE()
		//
	}
	r.GET("/validate", middleware.Userauthoraization, controllers.Validate)
	r.Run()

}
