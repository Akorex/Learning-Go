package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type TODO struct {
	ID   int
	Task string
}

var todos []TODO

func controller(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Access-Control-Allow-Origin", "*")

	query := r.URL.Query()

	id, _ := strconv.Atoi(query.Get("id"))
	fmt.Println(id)
}

func main() {

	http.HandleFunc("/", controller)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
