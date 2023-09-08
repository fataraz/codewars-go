package kyuu8

import (
	"fmt"
	"math"
)

//Write a method, that will get an integer array as parameter and will process every number from this array.
//
//Return a new array with processing every number of the input-array like this:
//
//If the number has an integer square root, take this, otherwise square the number.
//
//Example
//[4,3,9,7,2,1] -> [2,9,3,49,4,1]
//Notes
//The input array will always contain only positive numbers, and will never be empty or null.
//
//var _ = Describe("Example Tests", func() {
//	tests := [][][]int{
//		{{4, 3, 9, 7, 2, 1},{2, 9, 3, 49, 4, 1}},
//		{{100, 101, 5, 5, 1, 1},{10, 10201, 25, 25, 1, 1}},
//		{{1, 2, 3, 4, 5, 6},{1, 4, 9, 2, 25, 36}},
//	}
//	It("should test that the solution returns the correct value", func() {
//		for _, t := range tests {
//			Expect(SquareOrSquareRoot(t[0])).To(Equal(t[1]))
//		}
//	})
//})

func SquareOrSquareRoot(arr []int) (results []int) {
	for _, v := range arr {
		sqrt := math.Sqrt(float64(v))
		fmt.Println(math.Floor(sqrt))
		if sqrt == math.Floor(sqrt) {
			results = append(results, int(sqrt))
		} else {
			results = append(results, v*v)
		}
	}
	return
}
