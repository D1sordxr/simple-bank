package vo

import (
	"encoding/json"
	"github.com/D1sordxr/simple-banking-system/internal/domain/shared/outbox/exceptions"
)

type MessagePayload struct {
	Payload string
}

func NewMessagePayload(data interface{}) (MessagePayload, error) {
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return MessagePayload{}, exceptions.MarshalFailed
	}
	return MessagePayload{Payload: string(payloadBytes)}, nil
}

func (m MessagePayload) IsValid() bool {
	var temp interface{}
	return json.Unmarshal([]byte(m.Payload), &temp) == nil
}

func (m MessagePayload) Unmarshal(target interface{}) error {
	if !m.IsValid() {
		return exceptions.InvalidPayload
	}
	return json.Unmarshal([]byte(m.Payload), target)
}

func (m MessagePayload) String() string {
	return m.Payload
}
