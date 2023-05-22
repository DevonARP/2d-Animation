package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '🏀'
		maxFrames = 1200
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()
	buf := make([]rune, 0, width*height)

	for i := 0; i < maxFrames; i++ {
		px = +vx
		py = +vy

		if px <= 0 || px >= width-1 {
			vx += -1
		}

		if py <= 0 || py >= height-1 {
			vy += -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true

		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				//fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(time.Second / 20)
	}
}
