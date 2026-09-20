package classes

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"slices"
	"solopg/internal/domain/card/attributes/stats"
)

// - Class - //
type Class struct {
	name           string           `yaml:"name"`
	playable       bool             `yaml:"playable"`
	bonus          []stats.Modifier `yaml:"bonus"`
	equipementName string           `yaml:"armor_set"`
}

type Template = struct {
	Name           string           `yaml:"name"`
	Playable       bool             `yaml:"playable"`
	Bonus          []stats.Modifier `yaml:"bonus"`
	EquipementName string           `yaml:"armor_set"`
}

func New(params Template) Class {
	return Class{
		name:           params.Name,
		playable:       params.Playable,
		bonus:          params.Bonus,
		equipementName: params.EquipementName,
	}
}

func (c Class) GetName() string {
	return c.name
}

func (c Class) IsPlayable() bool {
	return c.playable
}

func (c Class) GetBonus() []stats.Modifier {
	return c.bonus
}

func (c Class) GetEquipementName() string {
	return c.equipementName
}

func (c Class) GetHash() string {
	encoded, _ := json.Marshal(Template{
		Name:           c.GetName(),
		Playable:       c.IsPlayable(),
		Bonus:          c.GetBonus(),
		EquipementName: c.GetEquipementName(),
	})
	digest := sha256.Sum256(encoded)
	return hex.EncodeToString(digest[:])
}

func Assert(provided string) bool {
	return provided == "" || slices.Contains(ListNames(), provided)
}
