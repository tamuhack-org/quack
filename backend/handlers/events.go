package handlers

import (
	// "fmt"
	"io/ioutil"
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

	url := globalConfig.EventbriteUrl + "/users/me/owned_events/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Error building request.", 400)
	}

	query := req.URL.Query()
	query.Add("token", globalConfig.EventbriteToken)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error hitting the eventbrite API.", 400)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error parsing the eventbrite API response.", 400)
	}

	w.Write(body)
}
