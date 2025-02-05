package commands

type CreateClientCommand struct {
	FirstName  string           `json:"first_name" binding:"required"`
	LastName   string           `json:"last_name" binding:"required"`
	MiddleName string           `json:"middle_name" binding:"required"`
	Email      string           `json:"email" binding:"required"`
	Phones     []map[string]int `json:"phones" binding:"required"`
}

type UpdateClientCommand struct {
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	MiddleName string           `json:"middle_name"`
	Email      string           `json:"email"`
	Phones     []map[string]int `json:"phones"`
	Status     string           `json:"status"`
}
