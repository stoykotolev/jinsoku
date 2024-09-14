package screens

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GameSummary(window fyne.Window, score int) fyne.CanvasObject {
	gseText := canvas.NewText("Game session has ended.", color.White)
	gseText.TextSize = 64
	scoreText := canvas.NewText(fmt.Sprintf("Your score is %d", score), color.White)
	scoreText.TextSize = 64

	centeredText := container.NewVBox(
		container.NewCenter(gseText),
		layout.NewSpacer(),
		container.NewCenter(scoreText),
	)

	backButton := widget.NewButton("Back to Home", func() {
		window.SetContent(MainScreen(window))
	})

	newRoundButton := widget.NewButton("New Game", func() {
		StartGame(window)
	})

	centeredBtns := container.NewHBox(
		layout.NewSpacer(),
		backButton,
		newRoundButton,
		layout.NewSpacer(),
	)

	//TODO: Setup logic to also store the past scores and maybe order them from highest to lowest.
	return container.NewVBox(
		layout.NewSpacer(),
		centeredText,
		layout.NewSpacer(),
		centeredBtns,
	)
}
