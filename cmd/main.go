package main

import (
	"fmt"
	"goTodoProject/config"
	"goTodoProject/controller"
	"goTodoProject/repository"
	"net/http"
)

// var tr = repository.NewMemoryTodoRepository()
var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	ro.Routing()
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.WebPort), nil))
}
