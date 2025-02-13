package vo

import (
	"encoding/json"
	"github.com/D1sordxr/simple-bank/bank-services/internal/domain/shared/event/exceptions"
)

type EventPayload struct {
	Payload string
}

func NewEventPayload(data any) (EventPayload, error) {
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return EventPayload{}, exceptions.MarshalFailed
	}
	return EventPayload{Payload: string(payloadBytes)}, nil
}

func (e EventPayload) IsValid() bool {
	var temp interface{}
	return json.Unmarshal([]byte(e.Payload), &temp) == nil
}

func (e EventPayload) Unmarshal(target interface{}) error {
	if !e.IsValid() {
		return exceptions.InvalidPayload
	}
	return json.Unmarshal([]byte(e.Payload), target)
}

func (e EventPayload) String() string {
	return e.Payload
}
