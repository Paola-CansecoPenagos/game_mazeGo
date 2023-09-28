package maze

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (s *Scoring) Reset() {
	s.score = 0
	s.updateScoreLabel()
}

func (s *Scoring) IncreaseScore(points int) {
	s.score += points
	s.updateScoreLabel()
}

func (s *Scoring) updateScoreLabel() {
	s.scoreLabel.SetText("Puntuación: " + strconv.Itoa(s.score))
}

func NewScoring() *Scoring {
	scoreLabel := widget.NewLabel("Puntuación: 0")

	return &Scoring{
		score:      0,
		scoreLabel: scoreLabel,
	}
}

func (s *Scoring) CreateScoreUI() fyne.CanvasObject {
	return s.scoreLabel
}

type Scoring struct {
	score      int
	scoreLabel *widget.Label
}
