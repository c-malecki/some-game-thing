package core

type Player struct {
	Name      string
	Health    int
	Mana      int
	Latitude  int
	Longitude int
}

func (p *Player) New(name string) {
	p.Name = name
}
