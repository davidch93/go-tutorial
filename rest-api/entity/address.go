package entity

// Address is an entity that represent address model.
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
