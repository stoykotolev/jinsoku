package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Key Detection App")

	// The character to match
	lettersToMatch := "{"

	// Display the letter in the UI
	text := canvas.NewText(lettersToMatch, color.White)
	text.TextSize = 64

	content := container.NewCenter(text)
	myWindow.SetContent(content)

	// Set a keyboard event listener
	myWindow.Canvas().SetOnTypedRune(func(r rune) {
		typed := string(r)
		fmt.Println("Pressed symbol", typed)
	})

	// Show the window and start the app
	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()

}
