package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now().Format("2006-01-02 15:04:05")
        fmt.Fprintf(w, "The current time is %s", currentTime)
    })

    fmt.Println("Listening on :8080...")
    http.ListenAndServe(":8080", nil)
}
