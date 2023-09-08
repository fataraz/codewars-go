package kyuu8

import "strings"

//Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.
//
//Examples (input -> output)
//6, "I"     -> "IIIIII"
//5, "Hello" -> "HelloHelloHelloHelloHello"

func RepeatStr(repetitions int, value string) (result string) {
	for i := 0; i < repetitions; i++ {
		result += value
	}
	return result
}

func RepeatStrBest(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}
