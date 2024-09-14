package screens

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainScreen(window fyne.Window) fyne.CanvasObject {
	headingText := widget.NewLabel("Welcome to Jinsoku!")
	headingText.Alignment = fyne.TextAlignCenter
	subText := widget.NewLabel("The game where you learn your symbols and numbers")
	subText.Alignment = fyne.TextAlignCenter

	topText := container.NewCenter(
		container.NewVBox(
			headingText,
			layout.NewSpacer(),
			subText,
		),
	)

	// Create the buttons, stacked on top of each other
	button1 := widget.NewButton("Settings", func() {
		window.SetContent(ConfigurationScreen(window))
	})

	countDownSlice := [5]string{"5", "4", "3", "2", "1"}
	button2 := widget.NewButton("Start Game", func() {
		countdown := canvas.NewText("", color.White)
		countdown.TextSize = 64
		window.SetContent(container.NewCenter(countdown))
		for _, el := range countDownSlice {
			countdown.Text = el
			time.Sleep(time.Second)
			countdown.Refresh()
		}
		window.SetContent(GameScren(window))
	})

	// Create a vertical box (VBox) with buttons, centered
	buttons := container.NewVBox(
		button1,
		layout.NewSpacer(),
		button2,
	)

	// Center the buttons using a layout
	buttonContainer := container.New(layout.NewCenterLayout(), buttons)

	// Combine top text and buttons with spacers in between to push content to top and bottom
	content := container.NewVBox(
		topText,            // Top text centered
		layout.NewSpacer(), // Spacer to push buttons down
		buttonContainer,    // Buttons centered at the bottom
	)

	return content

}
