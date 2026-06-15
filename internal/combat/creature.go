package combat

import (
	"internal/creatures"
)

type ICreature interface {
	GetName() string
	GetCurrentHealth() int
	GetMaxHealth() int
	GetPersonality() creatures.Personality
}

type Creature struct {
	Name string
	Species creatures.Species
	Personality creatures.Personality
}

func (c Creature) String() string {
	return string(c.Name[0])
}

func (c Creature) GetIcon() string {
 	if len(c.Name) == 0 { return " " }
	return string(c.Name[0])
}
