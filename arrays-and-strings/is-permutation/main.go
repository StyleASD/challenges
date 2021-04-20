package main

import (
	"sort"
)

func main() {
}

/**
	Given two strings, write a method to decide if one is a permutation of the other.

Test Data

(test, pest) = false
(abc, cba)   = true
(abcd, cba) = false

I added the type and the Constructor function for NewChars so that I can implement the sort interface. This allows me to
make the algorithm O(N*logN) as we sort the runes first and then loop through checking they are all equal. 
*/

type Chars []rune
func NewChars(str string) Chars {
	return []rune(str)
}

func (c Chars) Len() int {
	return len(c)
}

func (c Chars) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Chars) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func IsPermutation(strA string, strB string) bool {
	if len(strA) != len(strB) {
		return false // If the strings aren't the same length they can't be a permutation
	}

	var rA = NewChars(strA)
	var rB = NewChars(strB)

	sort.Sort(rA)
	sort.Sort(rB)

	for i, _ := range rA {
		if rA[i] != rB[i] {
			return false
		}
	}

	return true
}
