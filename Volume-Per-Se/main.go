package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Transaction struct {
    ID        string    `json:"id"`
    Amount    float64  `json:"amount"`
    Timestamp time.Time `json:"timestamp"`
}

type Transactions []Transaction

func (t Transactions) CalculateVolumePerSecond() float64 {
    var totalVolume float64
    for _, transaction := range t {
        totalVolume += transaction.Amount
    }

    return totalVolume / time.Since(t[0].Timestamp).Seconds()
}

func main() {
    http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
        var transactions Transactions
        err := json.NewDecoder(r.Body).Decode(&transactions)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        volumePerSecond := transactions.CalculateVolumePerSecond()

        jsonResponse, err := json.Marshal(volumePerSecond)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
    })

    fmt.Println("Starting server on port 8080")
    http.ListenAndServe(":8080", nil)
}