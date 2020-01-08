package operator

import (
	"fmt"
)

func minOf(args ...int) int {
	min := args[0]

	for _, i := range args {
		if min > i {
			min = i
		}
	}

	return min
}

func LevenshteinDistance(str1 string, str2 string) [][]int {
	var str1len int = len(str1)
	var str2len int = len(str2)

	matrix := make([][]int, str1len+1)

	for i := 0; i < str1len+1; i++ {
		matrix[i] = make([]int, str2len+1)
		matrix[i][0] = i
	}

	matrix[0][0] = 0

	for j := 0; j < len(matrix[0]); j++ {
		matrix[0][j] = j
	}

	var cost int

	for i := 1; i < len(matrix); i++ {

		for j := 1; j < len(matrix[i]); j++ {
			if str1[i-1] == str2[j-1] {
				cost = 0
			} else {
				cost = 1
			}

			matrix[i][j] = minOf(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+cost)
		}
	}

	return matrix
}

func Print2D(matrix [][]int) {

	for _, elem := range matrix {
		for _, e := range elem {
			fmt.Print(fmt.Sprintf("|%d", e))
		}
		fmt.Println()
	}
}
