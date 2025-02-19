package tests

import (
	"context"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/client/commands"
	"github.com/D1sordxr/simple-bank/bank-services/internal/application/client/handlers"
	"testing"
)

func TestSuccessCreateClientHandler(t *testing.T) {
	command := commands.CreateClientCommand{
		FirstName:  "Oleg",
		LastName:   "Potapov",
		MiddleName: "Igorevich",
		Email:      "testing@mail.now",
		Phones: []map[string]int{
			{"country": 7, "code": 982, "number": 8823979},
			{"country": 1, "code": 555, "number": 1234567},
		},
	}

	ctx := context.Background()

	mockDeps := mockClientDeps()

	createClientService := handlers.NewCreateClientHandler(mockDeps)

	response, err := createClientService.Handle(ctx, command)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	expectedFullName := "Oleg Potapov Igorevich"
	if response.FullName != expectedFullName {
		t.Errorf("expected full name %q, got %q", expectedFullName, response.FullName)
	}

	expectedEmail := "testing@mail.now"
	if response.Email != expectedEmail {
		t.Errorf("expected email %q, got %q", expectedEmail, response.Email)
	}

	t.Logf("\n"+
		"Name: %s,\n"+
		"Email: %s,\n"+
		"Phones: %s\n"+
		"Status: %s\n"+
		"ClientID: %s\n",
		response.FullName, response.Email, response.Phones, response.Status, response.ClientID)
}
