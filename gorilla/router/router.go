// Package router manages route conttroller mapping for the server
package router

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/spf13/viper"

  "github.com/caddyshack/gorilla/controllers/core"
  "github.com/caddyshack/gorilla/controllers/note"
)

func Compile () (http.Handler) {
  p := viper.GetString("prefix")
  pfx := "/" + p + "/"
  r := mux.NewRouter()

  // Core
  r.HandleFunc(pfx,  core.HomeHandler).Methods("GET")

  // Note
  r.HandleFunc(pfx + "note/list", note.ListHandler).Methods("GET")
  r.HandleFunc(pfx + "note/{id:[0-9]+}", note.ReadHandler).Methods("GET")

  return r
}
