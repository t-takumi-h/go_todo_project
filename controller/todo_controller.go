package controller

import (
	"encoding/json"
	"fmt"
	"goTodoProject/controller/dto"
	"goTodoProject/entity"
	"goTodoProject/repository"
	"net/http"
	"github.com/oklog/ulid/v2"
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
	var todoResponses []dto.TodoResponse
	for _, v := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{Id: v.Id.String(), Title: v.Title, IsComplited: v.IsComplited})
	}
	
	todosResponse := dto.TodosResponse{Todos: todoResponses}

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

	todo := entity.TodoEntity{Title: todoRequest.Title, IsComplited: todoRequest.IsComplited}
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

	ulid, err := ulid.ParseStrict(todoRequest.Id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	todo := entity.TodoEntity{Id: ulid, Title: todoRequest.Title, IsComplited: todoRequest.IsComplited}
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

	ulid, err := ulid.ParseStrict(todoRequest.Id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	err = tc.tr.DeleteTodo(ulid)
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