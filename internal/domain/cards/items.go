package cards

type Item struct {
	Card

	Category Category
	Effects  []Effect
	Pod      int
}

type Category string

const (
	Weapon Category = "weapon"
	Armor  Category = "armor"
	Potion Category = "potion"
)

type Article struct {
	Item

	IsConsumable bool
}

type Wallet struct {
	Gold   int
	Silver int
	Copper int
}