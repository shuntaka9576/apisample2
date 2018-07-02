package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)
	router.GET("/todos/:todoId", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
