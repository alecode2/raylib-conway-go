package game

import rl "github.com/gen2brain/raylib-go/raylib"

func NewBoard(rows, cols int) Board {
	board := make(Board, rows)
	for x := range board {
		board[x] = make([]Tile, cols)

		for y := range board[x] {
			board[x][y].Alive = false
			board[x][y].Color = rl.Color{R: 0, G: 0, B: 0, A: 0}
			board[x][y].Age = uint8(255)
		}
	}

	return board
}

func ClearBoard(board Board) {
	for x := range board {
		for y := range board[x] {
			board[x][y].Alive = false
			board[x][y].Color = rl.Color{R: 0, G: 0, B: 0, A: 0}
			board[x][y].Age = uint8(255)
		}
	}
}

// TODO: Delete this in the near future, when we come up with a better testing solution
func AddShapes(board Board) {

	//glider
	board[1][1] = Tile{Alive: true, Color: rl.Blue}
	board[2][2] = Tile{Alive: true, Color: rl.Blue}
	board[3][2] = Tile{Alive: true, Color: rl.Blue}
	board[1][3] = Tile{Alive: true, Color: rl.Blue}
	board[2][3] = Tile{Alive: true, Color: rl.Blue}

	//blinker
	board[12][7] = Tile{Alive: true, Color: rl.Green}
	board[12][6] = Tile{Alive: true, Color: rl.Green}
	board[12][5] = Tile{Alive: true, Color: rl.Green}
}
