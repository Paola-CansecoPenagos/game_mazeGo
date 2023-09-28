package maze

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMovement(player *Player) *Movement {
	movement := &Movement{
		player:  player,
		buttons: make(map[string]*widget.Button),
	}

	upButton := widget.NewButton("↑", func() {
		movement.player.MovePlayer("up")
	})
	downButton := widget.NewButton("↓", func() {
		movement.player.MovePlayer("down")
	})
	leftButton := widget.NewButton("←", func() {
		movement.player.MovePlayer("left")
	})
	rightButton := widget.NewButton("→", func() {
		movement.player.MovePlayer("right")
	})

	movement.buttons["up"] = upButton
	movement.buttons["down"] = downButton
	movement.buttons["left"] = leftButton
	movement.buttons["right"] = rightButton

	return movement
}

func (m *Movement) CreateMovementUI() fyne.CanvasObject {
	// Aquí debes crear y devolver un objeto CanvasObject que represente la interfaz de movimiento.
	// Puedes usar un contenedor, como un HBox o VBox, para organizar los botones de dirección.
	return container.NewHBox(
		m.buttons["up"],
		m.buttons["down"],
		m.buttons["left"],
		m.buttons["right"],
	)
}

type Movement struct {
	player  *Player
	buttons map[string]*widget.Button
}
