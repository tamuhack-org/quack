package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// EmailSignupPost handles the email form submission
func EventsGET(w http.ResponseWriter, r *http.Request) {
	// Set header to form data, and validate a correct email.
	w.Header().Set("Content-Type", "application/json")

	log.Println("=========================")
	log.Println("Events handler!")
	fmt.Fprintf(w, "Events handler!")
	log.Println("=========================")
}
