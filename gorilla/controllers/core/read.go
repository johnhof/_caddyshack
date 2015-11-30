package core

import (
  "fmt"
  "net/http"
)

// HomeHandler returns a functionality mapping of the server
func HomeHandler (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello"))
}
