//You are given an M by N matrix consisting of booleans that represents a board. Each True boolean represents a wall. Each False boolean represents a tile you can walk on.
//
//Given this matrix, a start coordinate, and an end coordinate, return the minimum number of steps required to reach the end coordinate from the start. If there is no possible path, then return null. You can move up, left, down, and right. You cannot move through walls. You cannot wrap around the edges of the board.
//
//For example, given the following board:
//
//[[f, f, f, f],
//[t, t, f, t],
//[f, f, f, f],
//[f, f, f, f]]
//and start = (3, 0) (bottom left) and end = (0, 0) (top left), the minimum number of steps required to reach the end is 7, since we would need to go through (1, 2) because there is a wall everywhere else on the second row.

package main

import (
	"fmt"
	"math"
)

const (
	M = 9
	N = 10
	intmax = math.MaxInt64
)

type cell struct {
	x int
	y int
}

func minimumSteps(board [M][N]int, startcell cell, endcell cell) int64 {
	mem_board := make([][]int64, M)
	for i, _ := range mem_board {
		mem_board[i] = make([]int64, N)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			mem_board[i][j] = intmax
		}
	}
	mem_board[startcell.x][startcell.y] = 0
	findMinSteps(board, startcell, mem_board)
	return mem_board[endcell.x][endcell.y]
}

func findMinSteps(board [M][N]int, currentcell cell, mem_board [][]int64) {
	currentX := currentcell.x
	currentY := currentcell.y

	switch "up" {
	case "up":
		ok := validateCellPosition(currentX-1, currentY, board)
		if ok {
			updateAndLoop(currentcell, cell{x: currentX - 1, y: currentY}, mem_board, board)
		}
		fallthrough
	case "down":
		ok := validateCellPosition(currentX+1, currentY, board)
		if ok {
			updateAndLoop(currentcell, cell{x: currentX + 1, y: currentY}, mem_board, board)
		}
		fallthrough
	case "left":
		ok := validateCellPosition(currentX, currentY-1, board)
		if ok {
			updateAndLoop(currentcell, cell{x: currentX, y: currentY - 1}, mem_board, board)
		}
		fallthrough
	case "right":
		ok := validateCellPosition(currentX, currentY+1, board)
		if ok {
			updateAndLoop(currentcell, cell{x: currentX, y: currentY + 1}, mem_board, board)
		}
	}

	return
}

func updateAndLoop(currentcell cell, landingcell cell, mem_board [][]int64, board [M][N]int) {
	currentX := currentcell.x
	currentY := currentcell.y
	xpos := landingcell.x
	ypos := landingcell.y

	if mem_board[currentX][currentY]+1 < mem_board[xpos][ypos] {
		mem_board[xpos][ypos] = mem_board[currentX][currentY] + 1
		findMinSteps(board, landingcell, mem_board)
	} else {
		if mem_board[currentX][currentY] + 1 == mem_board[xpos][ypos] {
			findMinSteps(board, landingcell, mem_board)
		}
	}
}
func validateCellPosition(x, y int, board [M][N]int) bool {
	if x >= 0 && x < M && y >= 0 && y < N && board[x][y] != 0 {
		return true
	}
	return false
}

func main() {
	board := [M][N]int{
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 1, 1, 0, 1, 1, 1, 0, 1, 0},
		{1, 0, 1, 1, 1, 1, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 0, 0, 1},
	}
	startCell := cell{x: 0, y: 0}
	endCell := cell{x: 1, y: 1}
	minSteps := minimumSteps(board, startCell, endCell)
	if minSteps == intmax {
		fmt.Println("Shortest Path doesn't exist")
	} else {
		fmt.Println("Shortest path is", minSteps)
	}

}
