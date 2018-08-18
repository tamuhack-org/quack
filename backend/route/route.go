package route

import (
	"net/http"
	"os"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// Handlers for each route.
	"github.com/tamuhack-org/quack/backend/handlers"
)

// Load returns the routes and middleware
func Load() http.Handler {
	return routes()
}

func routes() http.Handler {
	// Define a new router.
	r := mux.NewRouter().StrictSlash(true)

	// Route for collecting emails on the show page!
	r.HandleFunc("/email-signup", handlers.EmailSignupPOST).Methods("POST")

	// For serving our static assets.
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	// Wrap the router with middleware for logging.
	return h.LoggingHandler(os.Stdout, r)
}
