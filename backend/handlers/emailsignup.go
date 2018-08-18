package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// EmailSignupPost handles the email form submission
func EmailSignupPOST(w http.ResponseWriter, r *http.Request) {
	if ok := validateEmail(r); !ok {
		http.Error(w, "", 400)
	}

	// Get email to update, and log it for now.
	content := r.FormValue("email")
	log.Println("----------------------------------------------")
	log.Println(content)
	log.Println("----------------------------------------------")

	// Success!
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Regex to validate that the input email is valid.
func validateEmail(r *http.Request) bool {
	fmt.Println(r.FormValue("email"))
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(r.FormValue("email"))
}
