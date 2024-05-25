package main

import (
    "fmt"
    "net/http"
)

type User struct {
    ID int
    Name string
}

func main() {
    // Create a new user.
    user := User{
        ID: 1,
        Name: "John Doe",
    }

    // Create a new HTTP handler.
    handler := func(w http.ResponseWriter, r *http.Request) {
        // Write the user data to the response.
        fmt.Fprintf(w, "%+v", user)
    }

    // Register the handler with the server.
    http.HandleFunc("/users", handler)

    // Start the server.
    http.ListenAndServe(":8080", nil)
}