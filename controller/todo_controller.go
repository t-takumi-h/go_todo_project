package controller

import (
	"encoding/json"
	"fmt"
	"goTodoProject/controller/dto"
	"goTodoProject/repository"
	"net/http"
)

// 外部パッケージに公開するインタフェース
type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	AddTodo(w http.ResponseWriter, r *http.Request)
	EditTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct{
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request){
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	todosResponse := dto.ComvertTodosResponse(todos)

	output, err := json.MarshalIndent(todosResponse, "", "\t")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(output)
	
}

func (tc *todoController) AddTodo(w http.ResponseWriter, r *http.Request){
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoResponse
	json.Unmarshal(body, &todoRequest)

	todo, err := todoRequest.ComvertTodoEntity()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	id, err := tc.tr.InsertTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	output, err := json.MarshalIndent(id, "", "\t")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(output)
}

func (tc *todoController) EditTodo(w http.ResponseWriter, r *http.Request){
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoResponse
	json.Unmarshal(body, &todoRequest)

	todo, err := todoRequest.ComvertTodoEntity()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	err = tc.tr.UpdateTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	output, err := json.MarshalIndent("ok", "", "\t")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(output)
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request){
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoResponse
	json.Unmarshal(body, &todoRequest)

	todo, err := todoRequest.ComvertTodoEntity()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	err = tc.tr.DeleteTodo(todo.Id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	result := struct{
		Result string `json:"result"`
	}{
		Result: "ok",
	}
	output, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(output)
}