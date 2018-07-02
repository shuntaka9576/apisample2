package main

type Todo struct {
	Name string `json:"name"`
	Completed string `json:"completed"`
	Due string `json:"due"`
	ID int `json:"id"`
}

type Todos []Todo
