package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/spf13/viper"

  "github.com/caddyshack/gorilla/router"
)

func main() {

  // Config settings
  viper.AddConfigPath("./configuration")
  viper.SetConfigName("index")
  err := viper.ReadInConfig()
  port := viper.GetString("port")
  host := viper.GetString("host")

  if err != nil {
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }

  // Load routes
  r := router.Compile()
  http.Handle("/", r)

  // Listen
  fmt.Printf("Listening at " + (host + ":" + port))
  err = http.ListenAndServe((host + ":" + port), nil)
  if err != nil {
    log.Fatal("Error: %s", err)
  }
}
