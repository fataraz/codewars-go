package kyuu8

//It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string.
//You're given one parameter, the original string. You don't have to worry with strings with less than two characters.

func RemoveChar(word string) string {
	result := ""
	if len(word) > 0 {
		result = word[1:]
		result = result[:len(result)-1]
	}
	return result
}

func RemoveCharBest(word string) string {
	return word[1 : len(word)-1]
}
