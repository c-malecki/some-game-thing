package main

import (
	"fmt"
	"game/core"
)

func main() {
	inv := core.Inventory{}
	inv.New(6, 4)

	torch := core.Item{
		Name:        "torch",
		Height:      2,
		Width:       1,
		Type:        0,
		Quantity:    1,
		Description: "Used to light up surrounding area",
	}

	inv.AddItem(&torch)
	fmt.Printf("%+v\n", inv.Items)
	for _, s := range inv.Grid {
		fmt.Printf("%+v\n", s)
	}
	inv.ListItems()
	inv.RemoveItem(0)
	fmt.Printf("%+v\n", inv.Items)
	for _, s := range inv.Grid {
		fmt.Printf("%+v\n", s)
	}
}
