package handlers

import (
	"LearningArch/internal/application/client/commands"
	"context"
	"testing"
)

func TestSuccessCreateClientHandler(t *testing.T) {
	command := commands.CreateClientCommand{
		FirstName:  "Oleg",
		LastName:   "Potapov",
		MiddleName: "Igorevich",
		Email:      "testing@mail.now",
	}

	ctx := context.Background()

	createClient := NewCreateClientHandler()
	response, err := createClient.Handle(ctx, command)
	if err != nil {
		t.Log("test failed")
		return
	}
	t.Logf("%s", response.FullName)
	t.Logf("%s", response.Email)
}
