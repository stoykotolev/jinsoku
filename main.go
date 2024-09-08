package main

import (
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
}

func (game *GameState) startGame(text *canvas.Text) {
	for {
		select {
		case inp := <-game.inputChan:
			if inp == game.selectedSymbol {
				game.cRound += 1
				newSymbol := getRandomSymbol()
				game.selectedSymbol = newSymbol
				text.Text = newSymbol
				text.Color = color.White
			} else {
				text.Color = color.RGBA{R: 255, B: 0, G: 0, A: 255}
			}
			if game.cRound > game.nRounds {
				text.Text = "Finito"
				text.Color = color.RGBA{R: 0, B: 0, G: 255, A: 255}
			}
			text.Refresh()
		}
	}
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
