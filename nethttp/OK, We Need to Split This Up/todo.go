package main

type Todo struct {
	Name string `json:"name"`
	Completed string `json:"completed"`
	Due string `json:"due"`
}

type Todos []Todo
