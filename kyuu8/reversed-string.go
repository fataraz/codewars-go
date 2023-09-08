package kyuu8

func Solution1(word string) (reverse string) {
	for _, c := range word {
		reverse = string(c) + reverse
	}
	return
}
