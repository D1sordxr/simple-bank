package vo

import "github.com/D1sordxr/simple-banking-system/internal/domain/shared/vo"

type Balance struct {
	AvailableMoney vo.Money
	FrozenMoney    vo.Money
}

func NewBalance() Balance {
	return Balance{
		AvailableMoney: vo.NewMoney(),
		FrozenMoney:    vo.NewMoney(),
	}
}

func (b *Balance) SubAvailableMoney(money vo.Money) {
	b.AvailableMoney.Sub(money)
}

func (b *Balance) AddFrozenMoney(money vo.Money) {
	b.FrozenMoney.Add(money)
}

func (b *Balance) EqAvailableMoney(money vo.Money) bool {
	if b.AvailableMoney.Eq(money) {
		return true
	}
	return false
}

func (b *Balance) AddAvailableMoney(money vo.Money) {
	b.AvailableMoney.Add(money)
}
func (b *Balance) SubFrozenMoney(money vo.Money) {
	b.FrozenMoney.Sub(money)
}
func (b *Balance) AddFrozenMoneyFromAvailable(money vo.Money) {
	b.AvailableMoney.Sub(money)
	b.FrozenMoney.Add(money)
}

func (b *Balance) DepositBalance(money vo.Money) Balance {
	b.AddAvailableMoney(money)
	return *b
}
func (b *Balance) Purchase(money vo.Money) Balance {
	b.AddFrozenMoneyFromAvailable(money)
	return *b
}
