package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
    var item Item
    json.NewDecoder(r.Body).Decode(&item)
    items = append(items, item)
    json.NewEncoder(w).Encode(item)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/items", getItems).Methods("GET")
    r.HandleFunc("/items", createItem).Methods("POST")

    http.ListenAndServe(":8000", r)
}
