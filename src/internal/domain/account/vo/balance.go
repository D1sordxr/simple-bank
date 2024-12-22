package vo

import "github.com/D1sordxr/simple-banking-system/internal/domain/shared/shared_vo"

type Balance struct {
	AvailableMoney shared_vo.Money
	FrozenMoney    shared_vo.Money
}

func NewBalance() Balance {
	return Balance{
		AvailableMoney: shared_vo.NewMoney(),
		FrozenMoney:    shared_vo.NewMoney(),
	}
}

func (b *Balance) SubAvailableMoney(money shared_vo.Money) {
	b.AvailableMoney.Sub(money)
}

func (b *Balance) AddFrozenMoney(money shared_vo.Money) {
	b.FrozenMoney.Add(money)
}

func (b *Balance) EqAvailableMoney(money shared_vo.Money) bool {
	if b.AvailableMoney.Eq(money) {
		return true
	}
	return false
}

func (b *Balance) AddAvailableMoney(money shared_vo.Money) {
	b.AvailableMoney.Add(money)
}
func (b *Balance) SubFrozenMoney(money shared_vo.Money) {
	b.FrozenMoney.Sub(money)
}
func (b *Balance) AddFrozenMoneyFromAvailable(money shared_vo.Money) {
	b.AvailableMoney.Sub(money)
	b.FrozenMoney.Add(money)
}

func (b *Balance) DepositBalance(money shared_vo.Money) Balance {
	b.AddAvailableMoney(money)
	return *b
}
func (b *Balance) Purchase(money shared_vo.Money) Balance {
	b.AddFrozenMoneyFromAvailable(money)
	return *b
}
