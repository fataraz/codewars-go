package kyuu8

import (
	"math"
)

//Given an array of integers.
//
//Return an array, where the first element is the count of positives numbers and the second element is sum of negative
//numbers. 0 is neither positive nor negative.
//
//If the input is an empty array or is null, return an empty array.
//
//Example
//For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65]

func CountPositivesSumNegatives(numbers []int) (arr []int) {
	var positive int
	var negative int
	for _, v := range numbers {
		if v > 0 {
			positive++
		}
		if math.Signbit(float64(v)) {
			negative += v
		}
	}
	arr = append(arr, positive)
	arr = append(arr, negative)
	return
}
