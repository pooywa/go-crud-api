package repository

import (
	"go-crud-api/internal/model"
)

var todos []model.Todo
var nextID = 1

func GetAll() []model.Todo {
	return todos
}

func GetByID(id int) *model.Todo {
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i]
		}
	}
	return nil
}

func Create(todo model.Todo) model.Todo {
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	return todo
}

func Update(id int, updated model.Todo) *model.Todo {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Title = updated.Title
			todos[i].Done = updated.Done
			return &todos[i]
		}
	}
	return nil
}

func Delete(id int) bool {
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
