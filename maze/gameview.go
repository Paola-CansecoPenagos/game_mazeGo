package maze

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowGameView() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Laberinto")

	game := NewGame()
	game.Start()

	gameView := game.GetGameView()

	backButton := widget.NewButton("Volver a Inicio", func() {

	})

	content := container.NewVBox(
		gameView,
		backButton,
	)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
