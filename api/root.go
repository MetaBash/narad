package api

import (
	"fmt"
	"net/http"
)

// # Root Route handles requests to the '/' route
func RootRoute(w http.ResponseWriter, r *http.Request) {
	// # Send a welcome message to the client
	fmt.Fprintf(w, "🚀 Welcome To Narad! 🚀")
}
