package core

import (
	"fmt"
	"slices"
)

type Inventory struct {
	Height int
	Width  int
	Grid   [][]int
	Items  []*Item
}

func (inv *Inventory) New(w int, h int) {
	inv.Width = w
	inv.Height = h
	grid := make([][]int, 0, h)
	for i := 0; i < h; i++ {
		row := make([]int, 0, w)
		for i := 0; i < w; i++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}
	inv.Grid = grid
}

func (inv *Inventory) CheckSpace(item *Item) [][]int {
	coords := make([][]int, 0, item.Height+item.Width)
	var w, h int
	for y := range inv.Height {
		if h == item.Height {
			break
		}
		w = 0
		for x := range inv.Width {
			if w == item.Width {
				break
			}
			if inv.Grid[y][x] == 0 {
				coords = append(coords, []int{x, y})
				w += 1
			}
		}
		h += 1
	}
	return coords
}

func (inv *Inventory) AddItem(item *Item) {
	coords := inv.CheckSpace(item)
	if fits := len(coords) == item.Height*item.Width; fits {
		for _, c := range coords {
			y := c[1]
			x := c[0]
			inv.Grid[y][x] = 1
		}
		item.InvCoords = coords
		fmt.Printf("Picked up %s.\n", item.Name)
		inv.Items = append(inv.Items, item)
	} else {
		fmt.Printf("Not enough inventory space for %s.\n", item.Name)
	}
}

func (inv *Inventory) RemoveItem(idx int) {
	item := inv.Items[idx]
	for _, c := range item.InvCoords {
		y := c[1]
		x := c[0]
		inv.Grid[y][x] = 0
	}
	inv.Items = slices.Delete(inv.Items, idx, 1)
	fmt.Printf("Removed %s\n", item.Name)
}

func (inv *Inventory) ListItems() {
	msg := "Inventory:\n"
	for _, item := range inv.Items {
		msg += fmt.Sprintf("%s\n", item.Name)
	}
	println(msg)
}
