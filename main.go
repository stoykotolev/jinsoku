package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"math/rand"
)

var SpecialSymbolsAndNumbers = []rune{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	'!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
	'-', '=', '[', ']', '\\', ';', '\'', ',', '.', '/',
	'_', '+', '{', '}', '|', ':', '"', '<', '>', '?',
	'`', '~',
}

func randomInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Key Detection App")
	arrLen := len(SpecialSymbolsAndNumbers)
	el := randomInRange(0, arrLen)

	// The character to match
	lettersToMatch := SpecialSymbolsAndNumbers[el]

	// Display the letter in the UI
	text := canvas.NewText(string(lettersToMatch), color.White)
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
