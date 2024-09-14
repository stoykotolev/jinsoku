package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/stoykotolev/jinsoku/internal/screens"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Multi-Screen App")

	// Set the initial screen to the main screen
	myWindow.SetContent(screens.MainScreen(myWindow))
	myWindow.Resize(fyne.NewSize(840, 680))

	// Show and run the application
	myWindow.ShowAndRun()

}
