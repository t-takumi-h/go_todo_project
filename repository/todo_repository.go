package repository

import (
	"fmt"
	"goTodoProject/entity"

	"github.com/oklog/ulid/v2"
)

type TodoRepository interface{
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id ulid.ULID, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
}

type memoryTodoRepository struct{
	todos []entity.TodoEntity
}

func NewMemoryTodoRepository() TodoRepository{
	var todos []entity.TodoEntity
	todos = append(todos, entity.TodoEntity{Id: entity.GenerateUlid(), Title: "test1-1", IsComplited: false})
	todos = append(todos, entity.TodoEntity{Id: entity.GenerateUlid(), Title: "test2-2", IsComplited: true})
	return &memoryTodoRepository{
		todos: todos,
	}
}

func (tr *memoryTodoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	return tr.todos, nil
} 

func (tr *memoryTodoRepository) InsertTodo(todo entity.TodoEntity) (id ulid.ULID, err error) {
	tr.todos = append(tr.todos, todo)
	return todo.Id, nil
}

func (tr *memoryTodoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	return nil
} 

func (tr *memoryTodoRepository) DeleteTodo(id int) (err error) {
	return nil
} 

type todoRepository struct{
}

func NewTodoRepository() TodoRepository{
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	cmd := fmt.Sprintf("SELECT id, title, is_complited FROM %s ORDER BY id DESC", TableNameTodos)
	rows, err := DbConnection.Query(cmd)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	
	todos = []entity.TodoEntity{}

	for rows.Next() {
		todo := entity.TodoEntity{}
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.IsComplited); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
} 

func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id ulid.ULID, err error) {
	ulid := entity.GenerateUlid()
	cmd := fmt.Sprintf("INSERT INTO %s (id, title, is_complited) VALUES (?, ?, ?)", TableNameTodos)
	_, err = DbConnection.Exec(cmd, ulid, todo.Title, todo.IsComplited)
	if err != nil {
		return ulid, err
	}
	return entity.GenerateUlid(), nil
}

func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	return nil
} 

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	return nil
} 