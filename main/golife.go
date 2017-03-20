package main

import (
	"math/rand"
	"time"
	"github.com/nsf/termbox-go"
)

const (
	backgroundColor = termbox.ColorBlack
	cellColor = termbox.ColorGreen
	drawInterval = 100 * time.Millisecond
	cellWidth = 2
	size int = 32
)

var (
	cells = [size][size]bool{}
	pause = true
)

func Reset(cells *[size][size]bool) {
	//MapCells(cells, fReset)
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[row]); col++ {
			cells[row][col] = rand.Intn(10) < 4
		}
	}
}

func Next(cells *[size][size]bool) (newCells [size][size]bool) {
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[row]); col++ {
			c := CountIfAlive(cells, row, col)
			if c == 3 {
				newCells[row][col] = true
			} else if c == 2 {
				newCells[row][col] = cells[row][col]
			} else {
				newCells[row][col] = false
			}
		}
	}
	return
}

func CountIfAlive(cells *[size][size]bool, row, col int) (count int) {
	if row > 0 {
		if col > 0 && cells[row - 1][col - 1] {
			count++
		}
		if cells[row - 1][col] {
			count++
		}
		if col < size - 1 && cells[row - 1][col + 1] {
			count++
		}
	}
	if col > 0 && cells[row][col - 1] {
		count++
	}
	if cells[row][col] {
		count++
	}
	if col < size - 1 && cells[row][col + 1] {
		count++
	}
	if row < size - 1 {
		if col > 0 && cells[row + 1][col - 1] {
			count++
		}
		if cells[row + 1][col] {
			count++
		}
		if col < size - 1 && cells[row + 1][col + 1] {
			count++
		}
	}
	return
}

func Draw(cells *[size][size]bool) {
	termbox.Clear(backgroundColor, backgroundColor)
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[row]); col++ {
			color := backgroundColor
			if cells[row][col] {
				color = cellColor
			}
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(col * cellWidth + i, row, ' ', color, color)
			}
		}
	}
	termbox.Flush()
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

	Reset(&cells)
	Draw(&cells)

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
					Reset(&cells)
					Draw(&cells)
				}
			}
		default:
			if !pause {
				termbox.Clear(backgroundColor, backgroundColor)
				cells = Next(&cells)
				Draw(&cells)
			}
		}
		time.Sleep(drawInterval)
	}
}
