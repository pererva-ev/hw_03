package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func genarateMatrix() (res [][]int) {
	length := 4
	max := 10

	for i := 0; i < length; i++ {
		tmp := []int{}
		for j := 0; j < length; j++ {
			tmp = append(tmp, rand.Intn(max))
		}
		res = append(res, tmp)
	}
	return
}

func TestPositiveSlice(t *testing.T) {
	matrix := genarateMatrix()
	fmt.Println("Matrix:")
	for i := range matrix {
		fmt.Println(matrix[i])
	}
	res := arrayFromMatrix(&matrix)
	fmt.Printf("\nArray: %v", res)
}
