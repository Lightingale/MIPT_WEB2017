package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type shortUrl struct {
	Address string `json:"key"`
}

type longUrl struct {
	Address string `json:"url"`	
}


var (
	urlStore = make(map[int]string)
	urlId = 0
)

func postUrl(w http.ResponseWriter, r *http.Request) {
	var source longUrl
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		panic(err)
	}
	urlStore[urlId] = source.Address
	short := shortUrl{strconv.Itoa(urlId)}
	urlId++	
	json.NewEncoder(w).Encode(&short)
}

func getUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl, err := strconv.Atoi(vars["key"])
	if err != nil {
		panic(err)
	}	
	if sourceUrl, ok := urlStore[shortUrl]; ok {
		http.Redirect(w, r, sourceUrl, 301)
	}
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", postUrl).Methods("POST")
	router.HandleFunc("/{key}", getUrl)
	http.ListenAndServe(":8082", router)
}