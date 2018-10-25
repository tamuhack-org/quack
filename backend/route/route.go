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

	// Route for registering for tamuhack.
  r.HandleFunc("/registration", handlers.RegistrationPOST).Methods("POST")

	// Route for collecting emails on the show page!
	r.HandleFunc("/email-signup", handlers.EmailSignupPOST).Methods("POST")

	// Route for returning all of our events.
	r.HandleFunc("/events", handlers.EventsGET).Methods("GET")

	// For serving our static assets.
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	// Wrap the router with middleware for logging.
	loggingHandler := h.LoggingHandler(os.Stdout, r)

	// For handling cross origin scripting (TODO(jaykhatri) this isn't secure)
	headersOk := h.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := h.AllowedOrigins([]string{"*"})
	methodsOk := h.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Wrap in a prebuild cors handler.
	return h.CORS(originsOk, headersOk, methodsOk)(loggingHandler)
}
