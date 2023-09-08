package kyuu8

//Kata Task
//I have a cat and a dog.
//
//I got them at the same time as kitten/puppy. That was humanYears years ago.
//
//Return their respective ages now as [humanYears,catYears,dogYears]
//
//NOTES:
//
//humanYears >= 1
//humanYears are whole numbers only
//Cat Years
//15 cat years for first year
//+9 cat years for second year
//+4 cat years for each year after that
//Dog Years
//15 dog years for first year
//+9 dog years for second year
//+5 dog years for each year after that

func CalculateYears(years int) [3]int {
	var cats int
	var dogs int
	for i := 1; i <= years; i++ {
		if i == 2 {
			cats += 9
			dogs += 9
			continue
		} else if i > 2 {
			cats += 4
			dogs += 5
			continue
		}
		cats += 15
		dogs += 15
	}
	return [3]int{years, cats, dogs}
}

func CalculateYears2(years int) (result [3]int) {
	switch years {
	case 1:
		result = [3]int{1, 15, 15}
	case 2:
		result = [3]int{2, 24, 24}
	default:
		result = [3]int{years, 24 + 4*(years-2), 24 + 5*(years-2)}
	}
	return
}