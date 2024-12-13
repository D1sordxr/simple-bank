package commands

import "github.com/google/uuid"

type CreateClientCommand struct {
	ClientID   uuid.UUID
	FirstName  string
	LastName   string
	MiddleName string
	Email      string
	Phones     []map[string]int
	Status     string
}

type UpdateClientCommand struct {
	ClientID   uuid.UUID
	FirstName  string
	LastName   string
	MiddleName string
	Email      string
	Phones     []map[string]int
	Status     string
}
