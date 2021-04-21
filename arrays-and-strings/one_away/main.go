package main

func main() {

}

/**
Attempt 1:

We can exit early if the length is more than 2 away because it would take more that 2 changes to reach the desired
string.

We are given 3 actions that we can perform each action should cost us 1 point. So we need to check if we only need to
perform a single action incuring the cost of 1 point. I needed to look this up as I had no idea how to write a Levenshtein Distance.

The algorthing apears to be an O(N) in space and an O(N^2) in time complexity.

I think there is probably a solution which is O(N) time complexity and O(1) space complexity but I will need to think about it.
Calculating the full Levenshtein distance is propably overkill for checking changes to a single character.

*/
func OneAway(first, second string) bool {
	if first == second {
		return true
	}

	// If the strings are more than one character away we can exit early
	if len(first)-len(second) > 1 || len(second)-len(first) > 1 {
		return false
	}

	var distance = calculateLevenshtein([]rune(first), []rune(second))

	return distance == 1
}

// The algorithm is based on what was found here https://www.golangprograms.com/golang-program-for-implementation-of-levenshtein-distance.html
func calculateLevenshtein(first, second []rune) int {
	firstLength := len(first)
	secondLength := len(second)
	column := make([]int, len(first)+1)

	for y := 1; y <= firstLength; y++ {
		column[y] = y
	}

	for x := 1; x <= secondLength; x++ {
		column[0] = x
		lastKey := x - 1
		for y := 1; y <= firstLength; y++ {
			oldKey := column[y]
			var incr int
			if first[y-1] != second[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastKey+incr)
			lastKey = oldKey
		}
	}
	return column[firstLength]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

/**
After reviewing the examples there is an O(N) solution that involves checking if it is a replace or an (insert / remove)
operation. I will try to work it out on the next pass of performing this algorithm.
*/
