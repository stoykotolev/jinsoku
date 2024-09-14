package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ConfigurationScreen(window fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("This is Screen 1"),
		widget.NewButton("Back to Main Screen", func() {
			window.SetContent(MainScreen(window))
		}),
	)
}
