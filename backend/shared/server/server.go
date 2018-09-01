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
func Run(handlers http.Handler) {
	// Global config from env vars.
	globalConfig := config.GlobalConfig

	addr := fmt.Sprintf(":%d", globalConfig.Port)
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), addr)

	// Start the HTTP listener, and catch.
	log.Fatal(http.ListenAndServe(addr, handlers))
}
