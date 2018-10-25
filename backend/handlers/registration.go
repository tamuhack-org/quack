package handlers

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/tamuhack-org/quack/backend/shared/database"
)

// EmailSignupPost handles the email form submission
func RegistrationPOST(w http.ResponseWriter, r *http.Request) {
	// Grab the collection that holds all of our sign up emails.
	log.Println("enter")

	signups := database.Mongo.DB("").C("applicants")

	// Set header to form data, and validate a correct email.
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	// Get email to update, and write to DB.
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	log.Println("=========================")
	log.Println("Got " + firstName + " " + lastName + " from frontend.")
	log.Println("=========================")

	if _, err := signups.Upsert(bson.M{"first_name": firstName}, bson.M{"$set": bson.M{"last_name": lastName}}); err != nil {
		http.Error(w, "Error writing data to the database.", 400)
		return
	}

	// Log for debugging purposes.
	log.Println("=========================")
	log.Println("Wrote " + firstName + " " + lastName + " to database.")
	log.Println("=========================")
}
