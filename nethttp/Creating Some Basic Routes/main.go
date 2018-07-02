package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Print(ps.ByName("todoId"))
	fmt.Fprintf(w, "Todo show:%s", ps.ByName("todoId"))
}

func TodoNeko(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Print(ps.ByName("neko"))
	fmt.Fprintf(w, "Neko show:%s", ps.ByName("neko"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)
	router.GET("/todos/:todoId", TodoShow)
	router.GET("/t/:neko", TodoNeko)

	log.Fatal(http.ListenAndServe(":8080", router))
}
