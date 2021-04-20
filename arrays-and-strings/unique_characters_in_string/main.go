package main

/**
Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data
structures?

Test cases
- aadvark  = false
- and      = true
- Zoe      = true
- ZEAzPlp  = true
- /abcfts/ = false

This is the first implementation. However this uses an additional data structure a map to check if we have already seen
the character. This implementation is O(N) in time complexity and O(1) in space complexity. I dont think I can get
quicker in time complexity so I will see if I can implement and O(N) without the additional data strucutre.

To make the check without the 
*/
func StringHasUniqueCharacters_AttemptOne(str string) bool {
	data := map[rune]bool{}

	for _, char := range str {
		_, exists := data[char]

		if exists {
			return false
		}

		data[char] = true
	}

	return true
}
