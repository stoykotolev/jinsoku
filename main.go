package main

import (
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

func getRandomSymbol() string {
	arrLen := len(SpecialSymbolsAndNumbers)
	el := randomInRange(0, arrLen)

	// The character to match
	return string(SpecialSymbolsAndNumbers[el])
}

type Game struct {
	nRounds       int8
	currentSymbol string
}

func (g *Game ) gameLoop {
		if typed == g.currentSymbol {
			newSymbol := getRandomSymbol()
			game.currentSymbol = newSymbol
			text.Text = newSymbol
			text.Color = color.White
			text.Refresh()
		} else {
			text.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			text.Refresh()
		}
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Key Detection App")
	// Display the letter in the UI
	game := Game{
		nRounds:       10,
		currentSymbol: getRandomSymbol(),
	}
	text := canvas.NewText(game.currentSymbol, color.White)
	text.TextSize = 64

	content := container.NewCenter(text)
	myWindow.SetContent(content)

	// Start the game
	// Iterate N number of rounds, for each game

	// Set a keyboard event listener
	myWindow.Canvas().SetOnTypedRune(func(r rune) {
		typed := string(r)
	})

	// Show the window and start the app
	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()
}
