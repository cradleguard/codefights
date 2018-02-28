package main

func matrixElementsSum(matrix [][]int) int {
	leni := len(matrix)
	lenj := len(matrix[0])
	sum := 0
	for i := 0; i <= leni-1; i++ {
		for j := 0; j <= lenj-1; j++ {
			sum += matrix[i][j]
			if matrix[i][j] == 0 {
				for k := i + 1; k <= leni-1; k++ {
					matrix[k][j] = 0
				}
			}
		}
	}
	return sum
}
