package kyuu8

import (
	"sort"
	"strings"
)

//You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars)
//and then return the first value.
//
//The returned value must be a string, and have "***" between each of its letters.
//
//You should not remove or add elements from/to the array.

//s = []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
//Expect(TwoSort(s)).To(Equal("b***i***t***c***o***i***n"))
//s = []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
//Expect(TwoSort(s)).To(Equal("a***r***e"))
//s = []string{"lets", "talk", "about", "javascript", "the", "best", "language"}
//Expect(TwoSort(s)).To(Equal("a***b***o***u***t"))
//s = []string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}
//Expect(TwoSort(s)).To(Equal("c***o***d***e"))
//s = []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
//Expect(TwoSort(s)).To(Equal("L***e***t***s"))

func TwoSort(arr []string) (result string) {
	sort.Strings(arr)
	for _, v := range arr[0] {
		result += string(v) + "***"
	}
	return strings.TrimSuffix(result, "***")
}

// Best practice
func TwoSortBest(arr []string) string {
	sort.Strings(arr)
	return strings.Join(strings.Split(arr[0], ""), "***")
}
