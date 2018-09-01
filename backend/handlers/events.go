package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/globalsign/mgo/bson"
	"github.com/tamuhack-org/quack/backend/shared/database"
)

// EmailSignupPost handles the email form submission
func EmailSignupPOST(w http.ResponseWriter, r *http.Request) {
	// Set header to form data, and validate a correct email.
	w.Header().Set("Content-Type", "application/json")

	log.Println("=========================")
	log.Println("Wrote " + email + " to database.")
	log.Println("=========================")
}

