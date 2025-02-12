package queries

type GetByIDAccountDTO struct {
	Currency       string  `json:"currency"`
	AvailableMoney float64 `json:"balance"`
	FrozenMoney    float64 `json:"frozen_money"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
