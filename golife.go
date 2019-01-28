package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	backgroundColor = termbox.ColorBlack
	cellColor       = termbox.ColorGreen
	textColor       = termbox.ColorWhite
	cellWidth       = 2
	size            = 32
	marginTop       = 2
	marginLeft      = 4
)

var (
	cells        = [size][size]bool{}
	pause        = true
	drawInterval = 100 * time.Millisecond
)

func reset() {
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[row]); col++ {
			cells[row][col] = rand.Intn(10) < 4
		}
	}
}

func next() {
	newCells := [size][size]bool{}
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			c := countIfAlive(row, col)
			if c == 3 {
				newCells[row][col] = true
			} else if c == 2 {
				newCells[row][col] = cells[row][col]
			} else {
				newCells[row][col] = false
			}
		}
	}
	cells = newCells
}

func countIfAlive(row, col int) (count int) {
	if row > 0 {
		if col > 0 && cells[row-1][col-1] {
			count++
		}
		if cells[row-1][col] {
			count++
		}
		if col < size-1 && cells[row-1][col+1] {
			count++
		}
	}
	if col > 0 && cells[row][col-1] {
		count++
	}
	if col < size-1 && cells[row][col+1] {
		count++
	}
	if row < size-1 {
		if col > 0 && cells[row+1][col-1] {
			count++
		}
		if cells[row+1][col] {
			count++
		}
		if col < size-1 && cells[row+1][col+1] {
			count++
		}
	}
	return
}

func draw() {
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[row]); col++ {
			color := backgroundColor
			if cells[row][col] {
				color = cellColor
			}
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(marginLeft+col*cellWidth+i, marginTop+row, ' ', color, color)
			}
		}
	}
}

func drawText(x, y int, text string) {
	for _, c := range text {
		termbox.SetCell(size+4+x, y, c, textColor, backgroundColor)
		x++
	}
}

func drawKeyOperationsText() {
	y := marginTop
	drawText(size+8, y, "+----------+----------------+")
	y++
	drawText(size+8, y, "| Key      | Function       |")
	y++
	drawText(size+8, y, "+----------+----------------+")
	y++
	drawText(size+8, y, "| Space    | Run or Pause   |")
	y++
	drawText(size+8, y, "| Ctrl + R | Reset cells    |")
	y++
	drawText(size+8, y, "| Ctrl + U | Speed up       |")
	y++
	drawText(size+8, y, "| Ctrl + D | Speed down     |")
	y++
	drawText(size+8, y, "| Escape   | Exit           |")
	y++
	drawText(size+8, y, "+----------+----------------+")
	y++
}

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
	reset()
	draw()
	drawKeyOperationsText()
	termbox.Flush()

loop:
	for {
		select {
		case event := <-eventQueue:
			if event.Type == termbox.EventKey {
				if event.Key == termbox.KeyEsc {
					break loop
				} else if event.Key == termbox.KeySpace {
					pause = !pause
				} else if event.Key == termbox.KeyCtrlR {
					pause = true
					reset()
					draw()
				} else if event.Key == termbox.KeyCtrlD {
					drawInterval += 50 * time.Millisecond
				} else if event.Key == termbox.KeyCtrlU {
					if drawInterval > 50*time.Millisecond {
						drawInterval -= 50 * time.Millisecond
					}
				}
			}
		default:
			if !pause {
				termbox.Clear(backgroundColor, backgroundColor)
				next()
				draw()
				drawKeyOperationsText()
			}
		}
		termbox.Flush()
		time.Sleep(drawInterval)
	}
}
