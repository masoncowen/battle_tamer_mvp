package combat

import (
	"internal/creatures"
)

type ICreature interface {
	GetCurrentHealth() int
	GetMaxHealth() int
	GetPersonality() creatures.Personality
}

type Creature struct {
	Species creatures.Species
	Personality creatures.Personality
}

func (c Creature) String() string {
	return string(c.Species)
}

func (c Creature) GetIcon() string {
	return c.Species.GetIcon()
}
