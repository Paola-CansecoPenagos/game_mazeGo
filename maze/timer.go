package maze

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (t *Timer) Start() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	startTime := time.Now()
	endTime := startTime.Add(t.duration)

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()
			remainingTime := endTime.Sub(currentTime)
			if remainingTime <= 0 {
				t.timerLabel.SetText("Tiempo restante: 0s")
				t.done <- true
				return
			}
			t.timerLabel.SetText("Tiempo restante: " + remainingTime.String())
		case <-t.done:
			return
		}
	}
}

func NewTimer(duration time.Duration) *Timer {
	timerLabel := widget.NewLabel("Tiempo restante: " + duration.String())

	return &Timer{
		duration:   duration,
		timerLabel: timerLabel,
		done:       make(chan bool),
	}
}

func (t *Timer) CreateTimerUI() fyne.CanvasObject {
	// Aquí debes crear y devolver un objeto CanvasObject que represente la interfaz del temporizador.
	// Puedes usar un widget.Label u otro widget para mostrar el tiempo restante.
	return t.timerLabel
}

type Timer struct {
	duration   time.Duration
	timerLabel *widget.Label
	done       chan bool
}
