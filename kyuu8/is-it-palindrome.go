package kyuu8

import "strings"

//Write a function that checks if a given string (case insensitive) is a palindrome. A palindrome is a word, number, phrase,
//or other sequence of symbols that reads the same backwards as forwards, such as madam or racecar, the date and time 12/21/33 12:21,
//and the sentence: "A man, a plan, a canal – Panama".

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	n := len(str)

	for i := 0; i < n; i++ {
		n -= 1
		if str[i] != str[n] {
			return false
		}
	}
	return true
}
