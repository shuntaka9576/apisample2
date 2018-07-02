package main

import (
	"time"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
}

func main()  {
	router := httprouter.New()
	router.GET("/", TodoIndex)
	log.Fatal(http.ListenAndServe(":8080", router))
}
