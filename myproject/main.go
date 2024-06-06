package main

import (
	"fmt"
	"net/http"

	"github.com/srinucdacc/myproject/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
