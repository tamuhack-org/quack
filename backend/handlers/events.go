package handlers

import (
	"fmt"
	"log"
	"net/http"

	// Reading in environmental variables.
	"github.com/tamuhack-org/quack/backend/shared/config"
)

// EmailSignupPost handles the email form submission
func EventsGET(w http.ResponseWriter, r *http.Request) {
	// Set header to form data, and validate a correct email.
	w.Header().Set("Content-Type", "application/json")

	// Global config from env vars.
	globalConfig := config.GlobalConfig

	log.Println("=========================")
	log.Println("Events handler!")
	log.Println(globalConfig.EventbriteToken)
	log.Println(globalConfig.EventbriteUrl)
	log.Println("=========================")

	fmt.Fprintf(w, "Events handler!")
}
