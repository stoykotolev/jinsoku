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
	el := randomInRange(0, arrLen-1)

	// The character to match
	return string(SpecialSymbolsAndNumbers[el])
}

type GameState struct {
	nRounds        int
	selectedSymbol string
	cRound         int
	c              fyne.Canvas
	inputChan      chan string
	sessionScore   int
}

func (game *GameState) startGame(text *canvas.Text) {
	// Start the main game loop
gameSess:
	for {
		select {
		// Listen for the input
		case inp := <-game.inputChan:
			// if input is correct, increment the game round and update the symbol
			if inp == game.selectedSymbol {
				game.cRound += 1
				newSymbol := getRandomSymbol()
				game.selectedSymbol = newSymbol
				text.Text = newSymbol
				text.Color = color.White
				game.sessionScore += 200
				// if the symbol is incorrect, color red.
			} else {
				text.Color = color.RGBA{R: 255, B: 0, G: 0, A: 255}
			}
			// if the current round is more than the number of rounds, end the game session
			if game.cRound > game.nRounds {
				break gameSess
			}
			text.Refresh()
		}
	}
	game.c.SetContent(container.NewCenter(canvas.NewText(fmt.Sprintf("Game session is done. Your score is %d", game.sessionScore), color.White)))
	//TODO: Add 2 btns: Back to menu or Restart game session
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Key Detection App")
	game := GameState{
		nRounds:        5,
		selectedSymbol: getRandomSymbol(),
		cRound:         1,
		c:              myWindow.Canvas(),
		inputChan:      make(chan string),
		sessionScore:   0,
	}

	text := canvas.NewText(game.selectedSymbol, color.White)
	text.TextSize = 64

	content := container.NewCenter(text)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 300))
	go game.startGame(text)
	game.c.SetOnTypedRune(func(r rune) {
		game.inputChan <- string(r)
	})

	myWindow.ShowAndRun()
}
