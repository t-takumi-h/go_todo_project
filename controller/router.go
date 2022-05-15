package controller

import (
	"net/http"
	"os"
)

type Router interface {
	Routing()
}

type router struct{
	tc TodoController
}

func NewRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) Routing(){
	http.HandleFunc("/todos/", ro.handleTodosRequest)
}


func (ro *router) handleTodosRequest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	w.Header().Set("Content-Type", "application/json")
	
	prefix := "/todos/"

	switch r.URL.Path{
	case prefix + "get-todos":
		ro.tc.GetTodos(w, r)
	case prefix + "add-todo":
		ro.tc.PostTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}