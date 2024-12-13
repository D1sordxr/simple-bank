package commands

type CreateDTO struct {
	ClientID string `json:"client_id"`

	// Testing
	FullName string   `json:"full_name"`
	Email    string   `json:"email"`
	Phones   []string `json:"phones"`
	// Testing
}
