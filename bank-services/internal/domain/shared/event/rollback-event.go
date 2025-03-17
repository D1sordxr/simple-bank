package event

type Rollback struct {
	EventID string `json:"event_id"`
}

type RollbackEvents []Rollback
