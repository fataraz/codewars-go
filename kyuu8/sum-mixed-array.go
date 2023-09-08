package kyuu8

import (
	"fmt"
	"strconv"
)

//Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.
//
//Return your answer as a number.
//
//var _ = Describe("Test Example", func() {
//	It("Basic tests", func() {
//		doTest([]any{9, 3, "7", "3"}, 22)
//		doTest([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}, 42)
//		doTest([]any{"3", 6, 6, 0, "5", 8, 5, "6", 2,"0"}, 41)
//		doTest([]any{"1", "5", "8", 8, 9, 9, 2, "3"}, 45)
//		doTest([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}, 61)
//	})
//})

func SumMix(arr []any) (result int) {
	for _, v := range arr {
		switch v := v.(type) {
		case string:
			k, _ := strconv.Atoi(v)
			result += k
		case int:
			result += v
		}
	}
	return
}

// clever
func SumMixClever(arr []any) int {
	sum := 0
	for _, v := range arr {
		iv, _ := strconv.Atoi(fmt.Sprintf("%v", v))
		sum += iv
	}
	return sum
}
