package main

import (
	"github.com/mitsutoshi/golife/golife"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	backgroundColor = termbox.ColorBlack
	cellColor       = termbox.ColorGreen
	textColor       = termbox.ColorWhite
	cellWidth       = 2
	size            = 40
	marginTop       = 2
	marginLeft      = 4
	ptnTextStartY   = 14
)

var (
	drawInterval = 100 * time.Millisecond
	cells        = golife.NewCells(size)
	pattern      = map[rune][][]bool{
		'A': golife.Glider,
		'B': golife.GliderGun,
		'C': golife.Galaxy,
	}
	pause = true
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	termbox.Clear(backgroundColor, backgroundColor)
	cells.Reset()
	draw()
	drawKeyOperationsText()
	drawPatternText()
	termbox.Flush()

loop:
	for {
		select {
		case event := <-eventQueue:

			if event.Type == termbox.EventKey {

				if event.Key == termbox.KeyEsc {

					// Exit
					break loop

				} else if event.Key == termbox.KeySpace {

					// Pause cells
					pause = !pause

				} else if event.Key == termbox.KeyCtrlR {

					// Reset cells
					pause = true
					cells.Reset()
					draw()

				} else if event.Key == termbox.KeyArrowLeft {

					// Speed down
					drawInterval += 50 * time.Millisecond

				} else if event.Key == termbox.KeyArrowRight {

					// Speed up
					if drawInterval > 50*time.Millisecond {
						drawInterval -= 50 * time.Millisecond
					}

				} else if _, ok := pattern[event.Ch]; ok {

					// Plot cells pattern
					pause = true
					cells.Clear()
					cells.UpdateCells(pattern[event.Ch])
					draw()
				}
			}
		default:
			if !pause {
				termbox.Clear(backgroundColor, backgroundColor)
				cells.Next()
				draw()
				drawKeyOperationsText()
				drawPatternText()
			}
		}
		termbox.Flush()
		time.Sleep(drawInterval)
	}
}

func draw() {
	for row := 0; row < len(cells.Cells); row++ {
		for col := 0; col < len(cells.Cells[row]); col++ {
			color := backgroundColor
			if cells.Cells[row][col] {
				color = cellColor
			}
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(marginLeft+col*cellWidth+i, marginTop+row, ' ', color, color)
			}
		}
	}
}

func drawKeyOperationsText() {
	y := marginTop
	drawText(size+8, y, "Control")
	y++
	y++
	drawText(size+8, y, "Key       Function       ")
	y++
	drawText(size+8, y, "---------+----------------")
	y++
	drawText(size+8, y, "Space    | Run or Pause   ")
	y++
	drawText(size+8, y, "Ctrl + R | Reset cells    ")
	y++
	drawText(size+8, y, "->       | Speed up       ")
	y++
	drawText(size+8, y, "<-       | Speed down     ")
	y++
	drawText(size+8, y, "Escape   | Exit           ")
	y++
	y++
}

func drawPatternText() {
	y := ptnTextStartY
	drawText(size+8, y, "Pattern")
	y++
	y++
	drawText(size+8, y, "Key       | Name           ")
	y++
	drawText(size+8, y, "----------+----------------")
	y++
	drawText(size+8, y, "Shift + A | Glider         ")
	y++
	drawText(size+8, y, "Shift + B | Glider Gun     ")
	y++
	drawText(size+8, y, "Shift + C | Galaxy         ")
	y++
	y++
}

func drawText(x, y int, text string) {
	for _, c := range text {
		termbox.SetCell(size+4+x, y, c, textColor, backgroundColor)
		x++
	}
}
