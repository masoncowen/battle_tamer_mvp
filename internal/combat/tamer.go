package combat

type Tamer struct {
	Name string
}

type ITamer interface {
	GetName() string
}
