package controllers

import (
	"github.com/adeben33/golangMiniProject/Todo/initializers"
	"github.com/adeben33/golangMiniProject/Todo/models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetTodo(c *gin.Context) {
	var todo []models.Todo
	uniqueUser := c.MustGet("user")
	allTodo := initializers.DB.Model("models.Todo{}").Preload("Model.User").Where("todo.UserID <> ?", uniqueUser.(models.User).ID).Find(&todo)
	c.JSON(200, allTodo)
}
func GetTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	uniqueUser := c.MustGet("user")
	idTodo := initializers.DB.Model("models.Todo{}").
		Preload("Model.User").Where("todo.UserID <> ?", uniqueUser.(models.User).ID).
		Find(&todo)
	//idTodo := initializers.DB.Find(&todo, id)
	c.JSON(200, idTodo)
}

func CreateTodo(c *gin.Context) {
	type newTodo struct {
		content string
	}
	todo := c.bind(&newTodo{})
	user, err := c.Get("user")
	if err != nil {
		log.Fatal("Cant get the user")
	}

}
