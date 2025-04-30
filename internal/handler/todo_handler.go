package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go-crud-api/internal/model"
	"go-crud-api/internal/repository"
)

func RegisterTodoRoutes(r chi.Router) {
	r.Get("/", GetAllTodos)
	r.Get("/{id}", GetTodo)
	r.Post("/", CreateTodo)
	r.Put("/{id}", UpdateTodo)
	r.Delete("/{id}", DeleteTodo)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.GetAll())
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	todo := repository.GetByID(id)
	if todo == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	created := repository.Create(todo)
	json.NewEncoder(w).Encode(created)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var todo model.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	updated := repository.Update(id, todo)
	if updated == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if !repository.Delete(id) {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
