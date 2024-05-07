package core

type Player struct {
	Name   string
	Health int
	Mana   int
	Lat    int
	Long   int
	Items  []string
}

func (p *Player) New(name string) {
	p.Name = name
}
