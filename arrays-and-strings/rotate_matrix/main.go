package main

func main() {

}

/**
Attempt 1:

This ended up being an O(N^2) algorithm and takes up O(N) space.

So looking at the solution mine is fairly similar. I should have added a check to ensure the martix is N x N and this
can be done without without the additional space. As a result this can be done in place.

*/
func RotateMatrix_Attempt1(matrix [][]int) (rotated [][]int) {
	// Instantiate the rotated array
	rotated = make([][]int, len(matrix))

	for i, _ := range rotated {
		rotated[i] = make([]int, len(matrix))
	}

	// Loops to Perform the swaps
	var outerCounter int
	for i := len(matrix)-1; i > -1 ; i-- {
		currentArr := matrix[i];

		var innerCounter int
		for j := 0; j < len(currentArr); j++ {
			rotated[innerCounter][outerCounter] = currentArr[j]
			innerCounter++
		}

		outerCounter++
	}

	return rotated
}
