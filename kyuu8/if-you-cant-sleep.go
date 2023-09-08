package kyuu8

import "fmt"

//If you can't sleep, just count sheep!!
//
//Task:
//Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...".
//Input will always be valid, i.e. no negative integers.

func CountSheep(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		result += fmt.Sprintf("%d sheep...", i)
	}

	return result
}
