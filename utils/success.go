package utils

import (
	"encoding/json"
	"narad/schema"
	"net/http"
)

// # Send Success Response
func SendSuccess(w http.ResponseWriter, msg string) {
	// # Set Headers for Success Response
	w.WriteHeader(http.StatusOK) // # 200

	// # Create New Encoder for Success Body
	encoder := json.NewEncoder(w)

	// # Create Success Response Object
	res := &schema.Success{
		Success: true,
		Message: msg,
	}

	// # Encode Response to Writer
	encoder.Encode(res)
}

// # Send Payload Response
func SendPayload(w http.ResponseWriter, data interface{}) {
	// # Set Headers for Payload Response
	w.WriteHeader(http.StatusOK) // # 200

	// # Create New Encoder for Payload Body
	encoder := json.NewEncoder(w)

	// # Create Payload Response Object
	res := &schema.Payload{
		Success: true,
		Payload: data,
	}

	// # Encode Response to Writer
	encoder.Encode(res)
}
