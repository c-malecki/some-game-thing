package core

type GameStatus string

const (
	StatusStart GameStatus = "start"
	StatusMenu  GameStatus = "menu"
	StatusPlay  GameStatus = "play"
)

type GameState struct {
	Status GameStatus
	World  *MapGrid
}
