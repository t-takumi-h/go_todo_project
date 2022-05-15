package entity

import "github.com/oklog/ulid/v2"

type TodoEntity struct{
	Id ulid.ULID
	Title string
	IsComplited bool
}