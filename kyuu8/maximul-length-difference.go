package kyuu8

//You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z. Let x be any string in the first array and y be any string in the second array.
//
//Find max(abs(length(x) − length(y)))
//
//If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).
//
//Example:
//a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
//a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
//mxdiflg(a1, a2) --> 13
//Bash note:
//input : 2 strings with substrings separated by ,
//output: number as a string
//
//var _ = Describe("Tests MxDifLg", func() {
//
//	It("should handle basic cases", func() {
//		var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
//		var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
//		dotest(s1, s2, 13)
//		s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
//		s2 = []string{"bbbaaayddqbbrrrv"}
//		dotest(s1, s2, 10)
//	})
//})

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	min1, max1 := findMinAndMax(a1)
	min2, max2 := findMinAndMax(a2)

	diff1 := max2 - min1
	diff2 := max1 - min2

	if diff1 > diff2 {
		return diff1
	}
	return diff2
}

func findMinAndMax(strArr []string) (min int, max int) {
	min = len(strArr[0])
	max = len(strArr[0])
	for _, str := range strArr {
		strLength := len(str)
		if strLength > max {
			max = strLength
		} else if strLength < min {
			min = strLength
		}
	}
	return
}
