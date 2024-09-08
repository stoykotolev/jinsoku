package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

type GameState struct {
	nRounds       int
	currentSymbol string
	cRound        int
}

func (game *GameState) startGame(typed string, text *canvas.Text) {
	if typed == game.currentSymbol {
		newSymbol := getRandomSymbol()
		game.currentSymbol = newSymbol
		text.Text = newSymbol
		text.Color = color.White
		text.Refresh()
		game.cRound += 1
	} else {
		text.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
		text.Refresh()
	}
	if game.cRound > game.nRounds {
		text.Text = "Finito"
		text.Color = color.White
		text.Refresh()
	}
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Key Detection App")
	// Display the letter in the UI
	game := GameState{
		nRounds:       5,
		currentSymbol: getRandomSymbol(),
		cRound:        1,
	}

	text := canvas.NewText(game.currentSymbol, color.White)
	text.TextSize = 64

	content := container.NewCenter(text)
	myWindow.SetContent(content)

	// Start the game
	// Iterate N number of rounds, for each game
	// Keep score
	// Show total at the end
	myWindow.Resize(fyne.NewSize(300, 300))
	// Set a keyboard event listener
	myWindow.Canvas().SetOnTypedRune(func(r rune) {
		fmt.Println("Curr round", game.cRound, game.nRounds)
		typed := string(r)
		game.startGame(typed, text)
	})
	// Show the window and start the app
	myWindow.ShowAndRun()
}
