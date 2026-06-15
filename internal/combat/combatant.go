package combat

type Combatant struct { }

type ICombatant interface {
	GetName() string
	GetCurrentHealth() int
	GetMaxHealth() int
}
