package core

import (
	"fmt"
	"os"
)

func PrintGameMenu() {
	println(LineBreak)
	fmt.Printf("Menu Options:\n[1] help\n[2] save game\n[3] load game\n[4] quit game\n[5] exit menu\n")
}

func printHelp() {
	println(LineBreak)
	println("help menu")
}

func MenuCmd(cmd string, GS *GameState) {
	switch cmd {
	case "1":
		printHelp()
	case "2":
		PrintGameSlots()
	case "3":
		PrintGameSlots()
	case "4":
		println(LineBreak)
		println("good bye")
		os.Exit(0)
	case "5":
		GS.Status = "play"
	default:
		println("\ninvalid input\n")
		PrintGameMenu()
	}
}
