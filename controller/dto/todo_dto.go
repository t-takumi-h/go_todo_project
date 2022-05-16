package dto

import (
	"goTodoProject/entity"

	"github.com/oklog/ulid/v2"
)

type TodoResponse struct {
	Id string `json:"id"`
	Title string `json:"title"`
	IsComplited bool `json:"is_complited"`
}

func (todoResponse *TodoResponse) ComvertTodoEntity() (todoEntity entity.TodoEntity, err error){
	var id ulid.ULID
	if todoResponse.Id != "" {
		id, err = ulid.ParseStrict(todoResponse.Id)
		if err != nil {
			return entity.TodoEntity{}, err
		}
	}
	return entity.TodoEntity{Id: id, Title: todoResponse.Title, IsComplited: todoResponse.IsComplited}, nil
}


type TodosResponse struct{
	Todos []TodoResponse `json:"todos"`
}

func ComvertTodosResponse(todos []entity.TodoEntity) TodosResponse{
	var todoResponses []TodoResponse
	for _, v := range todos {
		todoResponses = append(todoResponses, TodoResponse{Id: v.Id.String(), Title: v.Title, IsComplited: v.IsComplited})
	}
	return TodosResponse{Todos: todoResponses}
}

