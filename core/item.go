package core

type Item struct {
	Name        string
	Height      int
	Width       int
	Type        int
	Quantity    int
	Description string
	InvCoords   [][]int
}

// ex
// Name torch
// Height 3 (unit like grid inventory?)
// Width 1 (unit like grid inventory?)
// Type environmental? (affects surrounding area for utility?)

func (i *Item) UseItem() {
	switch i.Type {
	case 0:
		// something
	}
}
