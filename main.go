package main

import (
	"fmt"
)

func isSquareMatrix(m [][]int) bool {
	l := len(m)
	for _, f := range m {
		if len(f) != l {
			return false
		}
	}

	return true
}

func arrayFromMatrix(matrix *[][]int) (res []int) {
	lenght := len(*matrix)
	x, y := -1, 0
	add := lenght
	dir := 1
	for add != 0 {
		for i := 0; i != add; i++ {
			x = x + dir
			res = append(res, (*matrix)[y][x])
		}
		add--
		for i := 0; i != add; i++ {
			y = y + dir
			res = append(res, (*matrix)[y][x])
		}
		dir = -dir
	}
	return
}

func main() {

	firstMatrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	secondMatrix := [][]int{{1, 2, 3, 2}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}}
	if !isSquareMatrix(firstMatrix) || !isSquareMatrix(secondMatrix) || (!isSquareMatrix(firstMatrix) && !isSquareMatrix(secondMatrix)) {
		fmt.Println("Number of columns and rows must be equal.")
		return
	}
	fmt.Println(arrayFromMatrix(&firstMatrix))
	fmt.Println(arrayFromMatrix(&secondMatrix))
}
