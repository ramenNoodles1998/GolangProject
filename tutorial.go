package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "github.com/gorilla/mux"
)

type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles map[string]Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the home page")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: singleArticle")
	vars:= mux.Vars(r)
	key := vars["id"]
	fmt.Println(Articles[key])
	json.NewEncoder(w).Encode(Articles[key])

}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("rest api v2.0 mux routers")
	Articles = map[string]Article {
		"1": {Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		"2": {Id: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	} 	
	handleRequests()
}