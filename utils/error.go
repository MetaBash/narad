package utils

import (
	"encoding/json"
	"narad/schema"
	"net/http"
)

// # Send Error Response
func SendError(w http.ResponseWriter, err error) {
	// # Set Headers for Error Response
	w.WriteHeader(http.StatusBadRequest) // # 400

	// # Create New Encoder for Error Body
	encoder := json.NewEncoder(w)

	// # Create Error Response Object
	res := &schema.Error{
		Success: false,
		Error:   err.Error(),
	}

	// # Encode Response to Writer
	encoder.Encode(res)
}
