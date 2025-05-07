package game

import rl "github.com/gen2brain/raylib-go/raylib"

func ConwayStep(gs *GameState) {
	ClearBoard(gs.Next)

	for x := range gs.Current {
		for y := range gs.Current[x] {
			tile := gs.Current[x][y]
			neighbors := MooreNeighbors(gs.Current, x, y)

			//We pass the color from one board to the other
			gs.Next[x][y].Color = gs.Current[x][y].Color

			if neighbors < 2 && tile.Alive {
				gs.Next[x][y].Alive = false
				gs.Next[x][y].Age = 1
			} else if neighbors >= 2 && neighbors <= 3 && tile.Alive {
				gs.Next[x][y].Alive = true
				gs.Next[x][y].Age = 0 // Reset age if cell is alive
			} else if neighbors > 3 && tile.Alive {
				gs.Next[x][y].Alive = false
				gs.Next[x][y].Age = 1
			} else if neighbors == 3 && !tile.Alive {
				// New cell born, blend color from neighbors
				gs.Next[x][y].Alive = true
				gs.Next[x][y].Color = BlendNeighborColors(gs.Current, x, y)
				gs.Next[x][y].Age = 0
			} else if !tile.Alive && tile.Age < 255 {
				gs.Next[x][y].Age = tile.Age + 1
			}
		}
	}
}

func MooreNeighbors(board Board, x, y int) (count int) {
	//North
	if y > 0 && board[x][y-1].Alive {
		count++
	}
	//South
	if y < len(board[x])-1 && board[x][y+1].Alive {
		count++
	}
	//West
	if x > 0 && board[x-1][y].Alive {
		count++
	}
	//East
	if x < len(board)-1 && board[x+1][y].Alive {
		count++
	}
	//NorthWest
	if x > 0 && y > 0 && board[x-1][y-1].Alive {
		count++
	}
	//NorthEast
	if x < len(board)-1 && y > 0 && board[x+1][y-1].Alive {
		count++
	}
	//SouthWest
	if x > 0 && y < len(board[x])-1 && board[x-1][y+1].Alive {
		count++
	}
	//SouthEast
	if x < len(board)-1 && y < len(board[x])-1 && board[x+1][y+1].Alive {
		count++
	}
	return count
}

func BlendNeighborColors(board Board, x, y int) rl.Color {
	count := 0
	r := 0
	g := 0
	b := 0
	a := 255

	//North
	if y > 0 && board[x][y-1].Alive {
		r += int(board[x][y-1].Color.R)
		g += int(board[x][y-1].Color.G)
		b += int(board[x][y-1].Color.B)
		count++
	}
	//South
	if y < len(board[x])-1 && board[x][y+1].Alive {
		r += int(board[x][y+1].Color.R)
		g += int(board[x][y+1].Color.G)
		b += int(board[x][y+1].Color.B)
		count++
	}
	//West
	if x > 0 && board[x-1][y].Alive {
		r += int(board[x-1][y].Color.R)
		g += int(board[x-1][y].Color.G)
		b += int(board[x-1][y].Color.B)
		count++
	}
	//East
	if x < len(board)-1 && board[x+1][y].Alive {
		r += int(board[x+1][y].Color.R)
		g += int(board[x+1][y].Color.G)
		b += int(board[x+1][y].Color.B)
		count++
	}
	//NorthWest
	if x > 0 && y > 0 && board[x-1][y-1].Alive {
		r += int(board[x-1][y-1].Color.R)
		g += int(board[x-1][y-1].Color.G)
		b += int(board[x-1][y-1].Color.B)
		count++
	}
	//NorthEast
	if x < len(board)-1 && y > 0 && board[x+1][y-1].Alive {
		r += int(board[x+1][y-1].Color.R)
		g += int(board[x+1][y-1].Color.G)
		b += int(board[x+1][y-1].Color.B)
		count++
	}
	//SouthWest
	if x > 0 && y < len(board[x])-1 && board[x-1][y+1].Alive {
		r += int(board[x-1][y+1].Color.R)
		g += int(board[x-1][y+1].Color.G)
		b += int(board[x-1][y+1].Color.B)
		count++
	}
	//SouthEast
	if x < len(board)-1 && y < len(board[x])-1 && board[x+1][y+1].Alive {
		r += int(board[x+1][y+1].Color.R)
		g += int(board[x+1][y+1].Color.G)
		b += int(board[x+1][y+1].Color.B)
		count++
	}

	return rl.Color{R: uint8(r / count), G: uint8(g / count), B: uint8(b / count), A: uint8(a)}
}
