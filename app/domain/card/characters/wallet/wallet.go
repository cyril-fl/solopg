package wallet

type Wallet struct {
	Gold   int `yaml:"gold"`
	Silver int `yaml:"silver"`
	Copper int `yaml:"copper"`
}

func NewWallet(gold, silver, copper int) Wallet {
	return Wallet{
		Gold:   gold,
		Silver: silver,
		Copper: copper,
	}
}
