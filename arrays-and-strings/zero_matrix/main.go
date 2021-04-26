package main

func main() {

}

/**
Attempt 1:

I think this can be broken into two steps the first check to see where all the 0 values are and the second to Zero
all the values.

From the example the Idea is correct but this can be done using less space something like an O(N) space or even in O(1)
space.

*/
func ZeroMatrix_Attempt1(matrix [][]int) [][]int {
	var zeros = findZeroes(matrix)

	// If we dont find any zeros we can exit the function early
	if len(zeros) < 1 {
		return matrix
	}

	return zeroMatrix(matrix, zeros)
}

func findZeroes(matrix [][]int) (foundZeros [][]int) {
	for i, arr := range matrix {
		for j, val := range arr {
			if val == 0 {
				foundZeros = append(foundZeros, []int{i, j})
			}
		}
	}

	return foundZeros
}

func zeroMatrix(matrix [][]int, zeros [][]int) [][]int {
	for _, zeros := range zeros {
		var row = zeros[0]
		var column = zeros[0]

		// Change the row
		matrix[row] = make([]int, len(matrix[row]))

		for _, rows := range matrix {
			rows[column] = 0
		}
	}

	return matrix
}
