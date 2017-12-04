package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Url struct {
	Address string `json:"url"`;
}

var (
	urlStore = make(map[int]string)
	urlId = 0
)

func postUrl(w http.ResponseWriter, r *http.Request) {
	var sourceUrl Url
	err := json.NewDecoder(r.Body).Decode(&sourceUrl)
	if err != nil {
		panic(err)
	}
	urlStore[urlId] = sourceUrl.Address
	shortUrl := Url{strconv.Itoa(urlId)}
	urlId++
	json.NewEncoder(w).Encode(&shortUrl)
}

func getUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl, err := strconv.Atoi(vars["key"])
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, urlStore[shortUrl], 301)
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", postUrl).Methods("POST")
	router.HandleFunc("/{key}", getUrl)
	http.ListenAndServe(":8082", router)
}