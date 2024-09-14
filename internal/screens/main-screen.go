package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainScreen(window fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Go to Screen 1", func() {
			window.SetContent(ConfigurationScreen(window))
		}),
	)
}
