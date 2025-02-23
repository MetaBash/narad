package schema

// # Success Response Schema
type Success struct {
	Success bool        `json:"success"` // # true
	Message interface{} `json:"message"`
}

// # Payload Response Schema
type Payload struct {
	Success bool        `json:"success"` // # true
	Payload interface{} `json:"payload"`
}

// # Error Response Schema
type Error struct {
	Success bool        `json:"success"` // # false
	Error   interface{} `json:"error"`
}
