package kyuu8

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func EvenOrOddClever(number int) string {
	return []string{"Even", "Odd"}[number&1]

}
