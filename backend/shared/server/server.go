package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server stores the hostname and port number.
type Server struct {
	Hostname string `json:"Hostname"`
	HTTPPort int    `json:"HTTPPort"`
}

// Run starts the HTTP.
func Run(handlers http.Handler, hostname string, port int) {
	addr := httpAddress(hostname, port)
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), addr)

	// Start the HTTP listener, and catch.
	log.Fatal(http.ListenAndServe(addr, handlers))
}

// Calculates the standard string pased to net/http to listen.
func httpAddress(hostname string, port int) string {
	return hostname + ":" + fmt.Sprintf("%d", port)
}
