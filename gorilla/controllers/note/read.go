package note

import (
  // "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

// ListHandler returns a list of notes
func ListHandler (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("List"))

}

// ReadHandler returns a single note
func ReadHandler (w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  w.Write([]byte("Read: " + id))
}
