package main

import "github.com/gin-gonic/gin"

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
