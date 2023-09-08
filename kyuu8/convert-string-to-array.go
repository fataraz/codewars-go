package kyuu8

import "strings"

//Write a function to split a string and convert it into an array of words.
//
//Examples (Input ==> Output):
//"Robin Singh" ==> ["Robin", "Singh"]
//
//"I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

func StringToArray(str string) (arr []string) {
	arr = strings.Split(str, " ")
	return
}
