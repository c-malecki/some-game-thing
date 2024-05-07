package core

import (
	"io/fs"
	"os"
)

func PrintStartScreen(dataFS fs.FS) {
	b, err := fs.ReadFile(dataFS, "data/assets/splash.txt")
	if err != nil {
		println(err)
	}
	println(string(b))
	// some like splash page thing
	start := "\n[1] New Game\n[2] Load Game\n[3] Options\n[4] About\n[5] Quit"
	println(start)
}

func StartNewGame() {
	// some intro text and image stuff
	// char creation
	println("Your name?")
}

func PrintGameSlots() {
	println("game slots")
}

func printOptions() {
	println("game options")
}

func printAbout() {
	println("about this game")
}

func StartCmd(cmd string, GS *GameState) {
	switch cmd {
	case "1":
		StartNewGame()
	case "2":
		PrintGameSlots()
	case "3":
		printOptions()
	case "4":
		printAbout()
	case "5":
		println(LineBreak)
		println("good bye")
		os.Exit(0)
	default:
		println("\ninvalid input\n")
		PrintGameMenu()
	}
}
