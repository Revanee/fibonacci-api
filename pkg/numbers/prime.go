package numbers

import "math"

func IsPrime(number int) bool {
	return trialDivision(number)
}

func trialDivision(number int) bool {
	if number <= 3 {
		return number > 1
	}
	if (number%2 == 0) || (number%3 == 0) {
		return false
	}

	count := 5
	for int(math.Pow(float64(count), 2)) <= number {
		if number%count == 0 || number%(count+2) == 0 {
			return false
		}
		count += 6
	}

	return true
}
