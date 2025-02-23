package router

import (
	"narad/api"
	"net/http"

	"github.com/gorilla/mux"
)

// # Router
func Router() *mux.Router {
	// # Create a new router object
	router := mux.NewRouter()

	// # Root Route
	router.HandleFunc("/", api.RootRoute).Methods("GET") // # Get server welcome message

	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("public"))))

	// # API Base Path
	BasePath := "/api/v1"

	// # API Endpoints Slice
	endpoints := []string{BasePath}

	// # Create a subrouter for the narad route
	narad := router.PathPrefix(endpoints[0]).Subrouter()

	// # Add routes to the subrouter
	narad.HandleFunc("/notification", api.SendNotification).Methods("POST")
	narad.HandleFunc("/live", api.GetLiveCard).Methods("GET")
	narad.HandleFunc("/history", api.GetHistoryCard).Methods("GET")
	narad.HandleFunc("/mark/{org_id}", api.MarkOrgInActive).Methods("PUT")
	narad.HandleFunc("/upload", api.UploadImage).Methods("POST")

	// # Return Router Object
	return router
}
