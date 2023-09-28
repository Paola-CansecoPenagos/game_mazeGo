package main

import (
	"laberintogo/maze"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Inicio del Laberinto")

	myApp.Settings().SetTheme(theme.LightTheme())

	startButton := widget.NewButton("Comenzar Juego", func() {
		go func() {
			game := maze.NewGame()
			game.Start()
		}()
	})

	content := container.NewVBox(
		startButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
