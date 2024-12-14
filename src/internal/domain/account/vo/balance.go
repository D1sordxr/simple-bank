package vo

type Balance struct {
	AvailableMoney Money
	FrozenMoney    Money
}

func NewBalance() Balance {
	return Balance{
		AvailableMoney: NewMoney(),
		FrozenMoney:    NewMoney(),
	}
}

func (b *Balance) SubAvailableMoney(money Money) {
	b.AvailableMoney.Sub(money)
}

func (b *Balance) AddFrozenMoney(money Money) {
	b.FrozenMoney.Add(money)
}

func (b *Balance) EqAvailableMoney(money Money) bool {
	if b.AvailableMoney.Eq(money) {
		return true
	}
	return false
}

func (b *Balance) AddAvailableMoney(money Money) {
	b.AvailableMoney.Add(money)
}
func (b *Balance) SubFrozenMoney(money Money) {
	b.FrozenMoney.Sub(money)
}
func (b *Balance) AddFrozenMoneyFromAvailable(money Money) {
	b.AvailableMoney.Sub(money)
	b.FrozenMoney.Add(money)
}

func (b *Balance) DepositBalance(money Money) Balance {
	b.AddAvailableMoney(money)
	return *b
}
func (b *Balance) Purchase(money Money) Balance {
	b.AddFrozenMoneyFromAvailable(money)
	return *b
}
