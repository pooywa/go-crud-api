package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "suit", Completed: false},
	{ID: "2", Item: "spaceship", Completed: false},
	{ID: "3", Item: "mars", Completed: false},
}

func main() {

	router := gin.Default()

	router.GET("/todos", getTodos)          //read
	router.GET("/todos/:id", getTodo)       //read
	router.POST("/todos", addTodo)          //create
	router.PUT("/todos/:id", updateTodo)    //update
	router.DELETE("/todos/:id", deleteTodo) //delete

	router.Run(`localhost:8080`)

}

func addTodo(context *gin.Context) {
	var newTodo = []todo{}

	err := context.BindJSON(&newTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, newTodo...)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoIndexById(id)
	if err != nil {
		context.IndentedJSON(http.StatusOK, todos[todo])
		return
	}
	context.IndentedJSON(http.StatusOK, todos[todo])
}

func getTodoIndexById(id string) (int, error) {
	for i, t := range todos {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func updateTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoIndexById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	if err := context.BindJSON(todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")

	index, err := getTodoIndexById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "ToDo not found"})
		return
	}
	todos = append(todos[:index], todos[index+1:]...)

	context.JSON(http.StatusOK, gin.H{"message": "ToDodeleted"})
}
