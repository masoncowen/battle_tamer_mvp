package creatures

type Species int

const (
	NoSpecies Species = iota
	Person
	GenericWolf
	PlantCannon
)

func (s Species) GetIcon() string {
	switch s {
	case Person:
		return "P"
	case GenericWolf:
		return "W"
	case PlantCannon:
		return "C"
	}
	return " "
}

func (s Species) String() string {
	switch s {
	case Person:
		return "Person"
	case GenericWolf:
		return "Wolf or Dog"
	case PlantCannon:
		return "Plant Cannon"
	}
	return "Unknown"
}
