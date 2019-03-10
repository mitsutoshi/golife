package golife

import "math/rand"

// Cells
type Cells struct {
	Size  int
	Cells [][]bool
}

// Initialize Cells struct
func NewCells(size int) *Cells {
	var cells [][]bool
	for i := 0; i < size; i++ {
		var cols []bool
		for j := 0; j < size; j++ {
			cols = append(cols, false)
		}
		cells = append(cells, cols)
	}
	return &Cells{Size: size, Cells: cells}
}

// Update cells with random value
func (cells *Cells) Reset() {
	for row := 0; row < len(cells.Cells); row++ {
		for col := 0; col < len(cells.Cells[row]); col++ {
			cells.Cells[row][col] = rand.Intn(10) < 4
		}
	}
}

// Clear cells (Change all cells to false)
func (cells *Cells) Clear() {
	for row := 0; row < len(cells.Cells); row++ {
		for col := 0; col < len(cells.Cells[row]); col++ {
			cells.Cells[row][col] = false
		}
	}
}

// Move cells to the next generation
func (cells *Cells) Next() {
	var rows [][]bool
	for row := 0; row < cells.Size; row++ {
		var cols []bool
		for col := 0; col < cells.Size; col++ {
			c := cells.CountIfAlive(row, col)
			if c == 3 {
				cols = append(cols, true)
			} else if c == 2 {
				cols = append(cols, cells.Cells[row][col])
			} else {
				cols = append(cols, false)
			}
		}
		rows = append(rows, cols)
	}
	cells.Cells = rows
}

// Count the number of surviving cells
func (cells *Cells) CountIfAlive(row, col int) (count int) {
	if row > 0 {
		if col > 0 && cells.Cells[row-1][col-1] {
			count++
		}
		if cells.Cells[row-1][col] {
			count++
		}
		if col < cells.Size-1 && cells.Cells[row-1][col+1] {
			count++
		}
	}
	if col > 0 && cells.Cells[row][col-1] {
		count++
	}
	if col < cells.Size-1 && cells.Cells[row][col+1] {
		count++
	}
	if row < cells.Size-1 {
		if col > 0 && cells.Cells[row+1][col-1] {
			count++
		}
		if cells.Cells[row+1][col] {
			count++
		}
		if col < cells.Size-1 && cells.Cells[row+1][col+1] {
			count++
		}
	}
	return
}

// Update cells value to specified pattern
func (cells *Cells) UpdateCells(pattern [][]bool) {
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			cells.Cells[i][j] = pattern[i][j]
		}
	}
}
