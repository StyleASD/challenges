package main

import (
	"strings"
)

func main() {

}

/**
	Attempt 1

	I couldn't figure a way of handling this one. Turns out there is a trick to it and we can concatonate s1 together.
	This must then contain the value of s2 if it is a rotation.
 */
func IsRotation(s1, s2 string) bool {
	// We can exit early if the lengths aren't the same or the strings are equal to each other
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}

	// We can concatenate s1 together because if s2 is a rotation it mush be contained
	s1s1 := s1 + s1

	return isSubstring(s1s1, s2)
}

func isSubstring(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}
