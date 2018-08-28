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
	// Grab the collection that holds all of our sign up emails.
	signups := database.Mongo.DB("").C("email-signups")

	// Set header to form data, and validate a correct email.
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	if ok := validateEmail(r); !ok {
		http.Error(w, "Invalid email.", 400)
		return
	}

	// Get email to update, and write to DB.
	email := r.FormValue("email")
	if _, err := signups.Upsert(bson.M{"email": email}, bson.M{"$set": bson.M{"email": email}}); err != nil {
		http.Error(w, "Error writing email to the database.", 400)
		return
	}

	// Log for debugging purposes.
	log.Println("=========================")
	log.Println("Wrote " + email + " to database.")
	log.Println("=========================")
}

// Regex to validate that the input email is valid.
func validateEmail(r *http.Request) bool {
	fmt.Println(r.FormValue("email"))
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(r.FormValue("email"))
}
