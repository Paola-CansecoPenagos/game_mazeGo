package maze

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Game struct {
	myApp    fyne.App
	myWindow fyne.Window
	maze     *Maze
	player   *Player
	movement *Movement
	scoring  *Scoring
	timer    *Timer
}

func NewGame() *Game {
	myApp := app.New()
	myWindow := myApp.NewWindow("Laberinto")

	mazeInstance := NewMaze()
	player := NewPlayer(mazeInstance)
	movement := NewMovement(player)
	scoring := NewScoring()
	timer := NewTimer(5 * time.Minute)

	return &Game{
		myApp:    myApp,
		myWindow: myWindow,
		maze:     mazeInstance,
		player:   player,
		movement: movement,
		scoring:  scoring,
		timer:    timer,
	}
}

func (g *Game) Initialize(window fyne.Window) {
	g.myWindow = window
}

func (g *Game) Start() {
	var controlContainer fyne.CanvasObject

	g.maze.GenerateMaze()

	g.player.Initialize(g.maze.playerX, g.maze.playerY)
	g.scoring.Reset()
	g.timer.Start()

	gameView := g.maze.RenderMaze()

	movementUI := g.movement.CreateMovementUI()

	scoreLabel := g.scoring.CreateScoreUI()

	timerLabel := g.timer.CreateTimerUI()

	controlContainer = container.NewHBox(
		widget.NewButton("Reiniciar", func() {
			g.maze.GenerateMaze()
			g.player.Initialize(g.maze.playerX, g.maze.playerY)
			g.scoring.Reset()
			g.timer.Start()

			gameView := g.maze.RenderMaze()
			g.myWindow.SetContent(container.NewVBox(
				gameView,
				widget.NewLabel("Controles del juego"),
				controlContainer,
			))
		}),
		movementUI,
		scoreLabel,
		timerLabel,
	)

	g.myWindow.SetContent(container.NewVBox(
		gameView,
		widget.NewLabel("Controles del juego"),
		controlContainer,
	))

	g.myWindow.Resize(fyne.NewSize(800, 600))
	g.myWindow.Show()
}

func (g *Game) GetGameView() fyne.CanvasObject {
	return g.maze.RenderMaze()
}
