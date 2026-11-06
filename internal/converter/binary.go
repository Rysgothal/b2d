package converter

import "math"

func Binary(value string) int {
	ans := 0
	size := len(value) - 1

	for _, digit := range value {
		if digit == '0' {
			size -= 1
			continue
		}

		ans += int(math.Pow(2, float64(size)))
		size -= 1
	}

	return ans
}
