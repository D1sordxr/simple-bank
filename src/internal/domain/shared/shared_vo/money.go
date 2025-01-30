package shared_vo

import "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_exceptions"

type Money struct {
	Value float64
}

func NewMoney() Money {
	return Money{Value: 0}
}

func NewMoneyFromFloat(money float64) (Money, error) {
	if money <= 0 {
		return Money{}, shared_exceptions.InvalidMoney
	}
	return Money{Value: money}, nil
}

func (m *Money) Sub(money Money) {
	m.Value -= money.Value
}

func (m *Money) Add(money Money) {
	m.Value += money.Value
}

func (m *Money) Eq(money Money) bool {
	if m.Value == money.Value {
		return false
	}
	return true
}
