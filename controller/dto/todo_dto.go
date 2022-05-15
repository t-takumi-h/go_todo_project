package dto

type TodoResponse struct {
	Id string `json:"id"`
	Title string `json:"title"`
	IsComplited bool `json:"is_complited"`
}

type TodosResponse struct{
	Todos []TodoResponse `json:"todos"`
}