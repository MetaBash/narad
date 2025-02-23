package main

import (
	"fmt"
	"narad/env"
	"narad/router"
	"net/http"
	"os"
	"strings"
)

func main() {
	// # Load Env Variables
	env.LoadEnv()

	// # Get Port
	port := env.GetEnv("PORT", "3000")

	// # Build Server URL with Port
	serverURI := os.Getenv("SERVER_URI")
	serverURI = strings.Replace(serverURI, "{PORT}", port, 1)

	// # Get Router Instance from Router Package
	r := router.Router()

	// # Start Server
	fmt.Printf("🚀 Server is running...\n")
	fmt.Printf("🔗 Link : %s\n", serverURI)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		// # Start Server Error
		fmt.Printf("🚫 Start Server Error : %v\n", err)
	}
}
