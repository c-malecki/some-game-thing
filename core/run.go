package core

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n> ")
	text, _ := reader.ReadString('\n')
	return text
}

func processCmd(input string, GS *GameState) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		println("no command received.")
		return
	}
	cmd := strings.ToLower(tokens[0])

	switch GS.Status {
	case "start":
		StartCmd(cmd, GS)
	case "menu":
		MenuCmd(cmd, GS)
	case "play":
		switch cmd {
		case "menu":
			GS.Status = "menu"
			PrintGameMenu()
		default:
			println("\ninvalid input\n")
		}
	}
}

func RunGame(dataFS fs.FS) {
	GS := GameState{
		Status: "start",
	}
	PrintStartScreen(dataFS)
	for {
		cmd := userInput()
		processCmd(cmd, &GS)
	}
}
