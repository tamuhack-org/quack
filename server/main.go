package main

import (
  "os"
	"fmt"

	"net/http"

  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
  loggedRouter := handlers.LoggingHandler(os.Stdout, r)
  fmt.Println("Listening in port :8080\n...")
  http.ListenAndServe(":8080", loggedRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
  fmt.Println("Hello")
}
