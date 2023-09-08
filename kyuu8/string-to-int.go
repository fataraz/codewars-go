package kyuu8

import "strconv"

func BinToDec(bin string) int {
	i, _ := strconv.ParseInt(bin, 2, 64)
	return int(i)
}
