package screens

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/stoykotolev/jinsoku/pkg/utils"
)

func StartGame(window fyne.Window) {
	countDownSlice := [5]string{"5", "4", "3", "2", "1"}
	countdown := widget.NewLabel("")
	window.SetContent(container.NewCenter(countdown))
	for _, el := range countDownSlice {
		countdown.SetText(el)
		time.Sleep(time.Second)
	}
	window.SetContent(GameScreen(window))
}

type GameState struct {
	nRounds        int
	selectedSymbol string
	cRound         int
	c              fyne.Canvas
	inputChan      chan string
	text           *canvas.Text
	times          []time.Time
	w              fyne.Window
}

func (game *GameState) drawText(symb string) {
	game.text.Text = symb
	game.selectedSymbol = symb
	game.text.Color = color.White
	game.text.Refresh()
}

func (game *GameState) startGame() {
	game.times = append(game.times, time.Now())
	game.c.SetOnTypedRune(func(r rune) {
		game.inputChan <- string(r)
	})
gameSess:
	for {
		select {
		case inp := <-game.inputChan:
			if inp == game.selectedSymbol {
				game.times = append(game.times, time.Now())
				game.cRound += 1
				newSymbol := utils.GetRandomSymbol()
				game.drawText(newSymbol)
			} else {
				game.text.Color = color.RGBA{R: 255, B: 0, G: 0, A: 255}
			}
			if game.cRound > game.nRounds {
				break gameSess
			}
			game.text.Refresh()
		}
	}
	score := 0
	for i, time := range game.times {
		if i == 0 {
			continue
		}
		diff := time.Sub(game.times[i-1])
		ms := diff.Milliseconds()
		switch {
		case ms < 500:
			score += 250
		case ms < 750:
			score += 150
		default:
			score += 50
		}
	}
	//TODO: Do a nerd test on actual perf between for/range with larger structure.
	game.w.SetContent(GameSummary(game.w, score))
}

func GameScreen(window fyne.Window) fyne.CanvasObject {
	symb := utils.GetRandomSymbol()
	text := canvas.NewText(symb, color.White)
	text.TextSize = 64
	game := GameState{
		nRounds:        5,
		selectedSymbol: symb,
		cRound:         1,
		c:              window.Canvas(),
		inputChan:      make(chan string),
		text:           text,
		times:          []time.Time{},
		w:              window,
	}

	go game.startGame()
	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(text),
		layout.NewSpacer(),
	)
	return content
}
