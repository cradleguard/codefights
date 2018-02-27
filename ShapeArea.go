package main

func shapeArea(n int) int {
	area := 1
	for i := 2; i <= n; i++ {
		area += (i - 1) * 4
	}
	return area
}
