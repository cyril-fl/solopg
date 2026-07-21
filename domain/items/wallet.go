package items

type Wallet struct {
	Gold   int
	Silver int
	Copper int
}

func NewWallet(gold, silver, copper int) Wallet {
	return Wallet{
		Gold:   gold,
		Silver: silver,
		Copper: copper,
	}
}