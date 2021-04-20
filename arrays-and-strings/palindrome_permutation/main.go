package main

import (
	"strings"
)

func main() {

}

/**
Attempt 1:

There is a bruit force approach where we could create every permutation and then check if it s a palindrome but this
would be an O(N!). Very Expensive.

My next thoughts is looking closer at the rules that make up a palindrome.
- If a string is an even length all characters must have an even amount to be a palindrome
- If the string length is odd then a single character much be odd while all others are even to be a palindrome

This gives me something to work with I count all the characters into a map in O(N) time. Using the rules above I can
then check that the counts conform to the above rules

Assumptions:
I can ignore whitespace
I can ignore capitals


Input: Tact Coa
Output: True (permutations: "taco cat", "atco cta", etc)
*/
func IsPermutationOfPalindrome(str string) bool {
	charMap, charCount := buildCharMap([]rune(strings.ToLower(str)))

	var isEven = (charCount % 2) == 0 // Check if the string length is odd or even

	return isPalindrome(charMap, isEven)
}

func isPalindrome(charMap map[rune]int, even bool) bool {
	if even {
		// all chars need to be event
		for _, i := range charMap {
			if i%2 != 0 {
				return false
			}
		}
	} else {
		var oddCount = 0
		for _, i := range charMap {
			if i%2 != 0 {
				oddCount++;
			}
		}

		if oddCount != 1 {
			return false
		}
	}

	return true
}

func buildCharMap(str []rune) (map[rune]int, int) {
	var charMap = map[rune]int{}
	var charCount = 0

	for _, char := range str {
		// I am assuming we can ignore whitespace
		if char == ' ' {
			continue
		}

		charCount++
		count, ok := charMap[char]

		if ok {
			charMap[char] = count + 1
		} else {
			charMap[char] = 1
		}

	}

	return charMap, charCount
}
