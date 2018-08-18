package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Message struct {
	Text     string
	Received time.Time
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	// Get for testing.
	// r.HandleFunc("/", HomeHandler).Methods("GET")

	// Logins should only be accepted as a POST.
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	// Logins should only be accepted as a POST.
	r.HandleFunc("/test", TestHandler).Methods("GET")

	// serve index.html in the static/ directory.
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	// Middleware to add log data to Stdout.
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	fmt.Println("Listening in port :8080\n...")
	http.ListenAndServe(":8080", loggedRouter)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	msg := Message{}

	// Parse json into our message struct.
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		panic(err)
	}

	// Test message is populated with time.
	msg.Received = time.Now().Local()
	msgJson, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	// Make the response a little more legit.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msgJson)
	fmt.Println("Calculated data!")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("here's some data that needs to be encyrpted with sessions: chicken"))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	msg := &Message{Text: "Alex", Received: time.Now().Local()}

	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
