package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	Color rl.Color
	Alive bool
	Age   uint8
}

// board alias
type Board [][]Tile

type GameState struct {
	ScreenWidth  int32
	ScreenHeight int32
	FPS          int32
	SimFPS       int32

	IsPaused   bool
	IsMenuOpen bool

	StepForward   bool
	SelectedColor rl.Color

	Lapsed   float32
	Step     float32
	CellSize int32

	BoardA  Board
	BoardB  Board
	Current Board
	Next    Board
}

func (s *GameState) SwapBoards() {
	s.Current, s.Next = s.Next, s.Current
}

func (s *GameState) ToggleCell(x, y int32) {
	s.Current[x][y].Alive = !s.Current[x][y].Alive
	s.Current[x][y].Color = s.SelectedColor
}

func (s *GameState) ResetBoard() {
	s.BoardA = NewBoard(int(s.ScreenWidth)/int(s.CellSize), int(s.ScreenHeight)/int(s.CellSize))
	s.BoardB = NewBoard(int(s.ScreenWidth)/int(s.CellSize), int(s.ScreenHeight)/int(s.CellSize))
	s.Current = s.BoardA
	s.Next = s.BoardB
}

/*
For possible NewGameState function we could use this for clarity
	rows := int(screenHeight / cellSize)
	cols := int(screenWidth / cellSize)
*/
