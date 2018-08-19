package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// EmailSignupPost handles the email form submission
func EmailSignupPOST(w http.ResponseWriter, r *http.Request) {
  // NOTE: We must write the return header right away. 
  // i.e. This code block needs to happen first.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if ok := validateEmail(r); !ok {
		http.Error(w, "something went wrong", 400)
	}

	// Get email to update, and log it for now.
	content := r.FormValue("email")
	log.Println("----------------------------------------------")
	log.Println(content)
	log.Println("----------------------------------------------")
}

// Regex to validate that the input email is valid.
func validateEmail(r *http.Request) bool {
	fmt.Println(r.FormValue("email"))
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(r.FormValue("email"))
}
