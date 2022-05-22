package controller

import (
	"context"
	"fmt"
	"goTodoProject/middleware/auth"
	"net/http"
	"os"
	"text/template"

	"google.golang.org/api/idtoken"
)

type Router interface {
	Routing()
}

type router struct {
	tc TodoController
}

func NewRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) Routing() {
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/get-token", auth.GetTokenHandler)
	http.Handle("/todos/", auth.MiddlewareAuth(http.HandlerFunc(ro.handleTodosRequest)))

}

func (ro *router) handleTodosRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	w.Header().Set("Content-Type", "application/json")

	prefix := "/todos/"

	switch r.URL.Path {
	case prefix + "get-todos":
		ro.tc.GetTodos(w, r)
	case prefix + "add-todo":
		ro.tc.AddTodo(w, r)
	case prefix + "edit-todo":
		ro.tc.EditTodo(w, r)
	case prefix + "delete-todo":
		ro.tc.DeleteTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/test.html")
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	// fmt.Println(r.Form["credential"][0])
	payload, err := idtoken.Validate(context.Background(), r.Form["credential"][0], "539909250233-rtcchg7irrilghs3tugs97676rhmfg63.apps.googleusercontent.com")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(payload)

}
