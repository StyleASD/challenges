package main

func main() {
	Urlify([]rune("Mr John Smith"), 13)
}

/**
Attempt 1:

I am givent the true length of the string so we will call the string array as (K) and the true length as N. I think
can be done in N time. This implementation ends up with a space complexity of N as well as we create a new array equal to
the lenth of the origional sting array.

From the Examples this implemation can be made more space efficient by working backwards though the array. This lets us
write over the buffer without having to create a second array

examples

"Mr John Smith    ", 13 => "Mr%20John%20Smith"
*/
func Urlify(strArray []rune, trueLength int) []rune {
	var urlString = make([]rune, len(strArray))

	// Pos track the position in the urlString that we need to insert he current rune
	var pos = 0
	for i := 0; i < trueLength; i++ {
		var curr = strArray[i]

		if curr != ' ' {
			urlString[pos] = curr
			pos++
		} else {
			urlString[pos] = '%'
			urlString[pos+1] = '2'
			urlString[pos+2] = '0'
			pos += 3
		}

		i++
	}

	return urlString
}
