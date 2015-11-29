package main

import (
  "fmt"
  "net/http"
  "log"
)

func init () {
  http.HandleFunc("/", rootHandler)
  http.HandleFunc("/note", noteHandler)
}

func main() {
  addr := "localhost:8080"

  log.Printf("Listening on %s", addr);

  err := http.ListenAndServe(addr, nil)

  if err != nil {
    log.Fatal("Error: %s", err)
  }
}

func rootHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func noteHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Groceries, Bills, etc")
}
