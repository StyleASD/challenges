package main

import (
	"fmt"
	"strings"
)

func main() {

}

/**
Attempt 1

	I thing this can be done in O(N) time and am currently creating it in 0(N) space. This one wasn't too difficult to
	come up with a solution. I suspect there is a trick which I am unaware of though.

	The main O(*) saving we can get is in the space complexity.

	From the Examples
	We have to make sure we are efficiently building strings. If space is the constrain we can run a count function to
	work out how long the compression will be in O(N) time before we run the compression. This will allow us to check if
	the length will save space. We can also use this count to set the string builder buffer size

*/
func CompressString_Attempt1(str string) string {
	var compressed = doCompression_Attempt1([]rune(str))

	// Id we dont save space after the compression we return the original string
	if len(compressed) >= len(str) {
		return str
	}

	return compressed
}

func doCompression_Attempt1(runes []rune) string {
	var compressed strings.Builder

	// Set the first character and initialise the counter
	var lastChar rune = runes[0]
	var count int8 = 1

	// We can start the loop form the second character
	for _, currentChar := range runes[1:] {
		if lastChar == currentChar {
			count++
		} else {
			compressed.WriteRune(lastChar)
			compressed.Write(convertNumberToRune(count))

			// Reset the counter and change the last character
			count = 1
			lastChar = currentChar
		}
	}

	// After the loop we add the final character
	compressed.WriteRune(lastChar)
	compressed.Write(convertNumberToRune(count))

	return compressed.String()
}

func convertNumberToRune(i int8) []byte {
	return []byte(fmt.Sprintf("%d", i))
}
