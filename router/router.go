package router

import (
	"narad/api"

	"github.com/gorilla/mux"
)

// # Router
func Router() *mux.Router {
	// # Create a new router object
	router := mux.NewRouter()

	// # Root Route
	router.HandleFunc("/", api.RootRoute).Methods("GET") // # Get server welcome message

	// # API Base Path
	BasePath := "/api/v1"

	// # API Endpoints Slice
	endpoints := []string{BasePath}

	// # Create a subrouter for the narad route
	narad := router.PathPrefix(endpoints[0]).Subrouter()

	// # Add routes to the subrouter
	narad.HandleFunc("/notification", api.SendNotification).Methods("POST") // # Send the notification
	narad.HandleFunc("/live", api.GetLiveCard).Methods("GET")               // # Get all notifications
	narad.HandleFunc("/history", api.GetHistoryCard).Methods("GET")         // # Get a single notification

	// # Return Router Object
	return router
}
