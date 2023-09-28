package maze

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
)

type cell struct {
	x int
	y int
}

func (m *Maze) GenerateMaze() {
	for y := range m.cells {
		for x := range m.cells[y] {
			m.cells[y][x] = 1
		}
	}

	startX := rand.Intn(len(m.cells[0]))
	startY := rand.Intn(len(m.cells))
	m.cells[startY][startX] = 0
	stack := []cell{{x: startX, y: startY}}

	directions := []cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		neighbors := []cell{}
		for _, dir := range directions {
			nx, ny := current.x+dir.x*2, current.y+dir.y*2
			if nx >= 0 && nx < len(m.cells[0]) && ny >= 0 && ny < len(m.cells) {
				if m.cells[ny][nx] == 1 {
					neighbors = append(neighbors, cell{nx, ny})
				}
			}
		}

		if len(neighbors) > 0 {
			next := neighbors[rand.Intn(len(neighbors))]
			m.cells[current.y][current.x] = 0
			m.cells[next.y][next.x] = 0
			stack = append(stack, next)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	m.playerX, m.playerY = startX, startY
	m.destinationX, m.destinationY = len(m.cells[0])-1, len(m.cells)-1
}

func (m *Maze) RenderMaze() fyne.CanvasObject {
	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(len(m.cells[0])))

	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	for y := range m.cells {
		for x := range m.cells[y] {
			cellLabel := canvas.NewText(" ", white)
			if m.playerX == x && m.playerY == y {
				cellLabel.Text = "P"
			} else if m.destinationX == x && m.destinationY == y {
				cellLabel.Text = "D"
			} else if m.cells[y][x] == 1 {
				cellLabel.Text = "#"
			}

			grid.AddObject(cellLabel)
		}
	}

	return grid
}

func NewMaze() *Maze {

	cells := make([][]int, 10)
	for i := range cells {
		cells[i] = make([]int, 10)
	}
	maze := &Maze{
		cells: cells,
	}
	maze.GenerateMaze()

	return maze
}

type Maze struct {
	cells        [][]int
	playerX      int
	playerY      int
	destinationX int
	destinationY int
	mazeView     *canvas.Rectangle
}
