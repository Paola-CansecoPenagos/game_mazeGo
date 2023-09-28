package maze

import (
	"fyne.io/fyne/v2"
)

func (p *Player) Initialize(startX, startY int) {
	p.x, p.y = startX, startY
	p.maze.RenderMaze()
}

type Player struct {
	x, y int
	maze *Maze
	win  fyne.Window
}

func NewPlayer(maze *Maze) *Player {
	player := &Player{
		x:    maze.playerX,
		y:    maze.playerY,
		maze: maze,
	}
	return player
}

func (p *Player) MovePlayer(direction string) {
	switch direction {
	case "up":
		if p.y > 0 && p.maze.cells[p.y-1][p.x] == 0 {
			p.y--
		}
	case "down":
		if p.y < len(p.maze.cells)-1 && p.maze.cells[p.y+1][p.x] == 0 {
			p.y++
		}
	case "left":
		if p.x > 0 && p.maze.cells[p.y][p.x-1] == 0 {
			p.x--
		}
	case "right":
		if p.x < len(p.maze.cells[0])-1 && p.maze.cells[p.y][p.x+1] == 0 {
			p.x++
		}
	}
	p.maze.RenderMaze()
}
