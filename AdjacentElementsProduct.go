package main

func adjacentElementsProduct(inputArray []int) int {
	maxProduct := inputArray[0] * inputArray[1]
	for i := 0; i <= len(inputArray)-2; i++ {
		currentProduct := inputArray[i] * inputArray[i+1]
		if maxProduct < currentProduct {
			maxProduct = currentProduct
		}

	}
	return maxProduct
}
