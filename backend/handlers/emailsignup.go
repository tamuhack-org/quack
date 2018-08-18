package handlers

import (
	"log"
	"net/http"
	"regexp"
)

// NotepadCreatePOST handles the note creation form submission
func EmailSignupPOST(w http.ResponseWriter, r *http.Request) {
	if ok := validateEmail(r); !ok {
		http.Error(w, "Error validating POST", 400)
	}

	// Get email to update, and log it for now.
	content := r.FormValue("email")
	log.Println(content)

	// Success!
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Regex to validate that the input email is valid.
func validateEmail(r *http.Request) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(r.FormValue("email"))
}
