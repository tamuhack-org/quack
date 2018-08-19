package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// Reading in environmental variables.
	"github.com/tamuhack-org/quack/backend/shared/config"
)

// Run starts the HTTP server.
func Run(handlers http.Handler, config *config.Config ) {
	addr := httpAddress(config.Hostname, config.Port)
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), addr)

	// Start the HTTP listener, and catch.
	log.Fatal(http.ListenAndServe(addr, handlers))
}

// Calculates the standard string pased to net/http to listen.
func httpAddress(hostname string, port int) string {
	return hostname + ":" + fmt.Sprintf("%d", port)
}
