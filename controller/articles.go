package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Articles = []Article{
	Article{Id: "0", Title: "First title", Content: "Content 1"},
	Article{Id: "1", Title: "Second title", Content: "Content 2"},
}

func Create(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(body, &article)

	for _, item := range Articles {
		if article.Title == item.Title {
			fmt.Fprintln(io.MultiWriter(w), "Duplicate article title. Please fill another one")
			return
		}
	}

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
}

func Fetch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	fmt.Fprintln(w, "Unknown article")
}

func List(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			fmt.Fprintln(w, "Article with id: "+id+" has been removed")
		}
	}
}
