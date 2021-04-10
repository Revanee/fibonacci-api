package numbers

import (
	"math"
)

func IsPerfectSquare(number int) bool {
	// A number is perfect square if it's root is an integer
	root := int(math.Sqrt(float64(number)))
	return root*root == number
}

func IsFibonacci(number int) bool {
	// False if number is negative
	if number < 0 {
		return false
	}

	// if one of 5(x^2) + 4 or 5(x^2) - 4 is a perfect square, then the number is Fibonacci
	// source: Wikipedia
	a := 5*int(math.Pow(float64(number), 2)) + 4
	b := 5*int(math.Pow(float64(number), 2)) - 4
	return IsPerfectSquare(a) || IsPerfectSquare(b)
}

func PreviousFibonacci(number int) int {
	// Fail if not fibonacci
	if !IsFibonacci(number) {
		return -1
	}
	// 1 is ambiguous, assuming unique number
	if number == 1 {
		return 0
	}
	// Divide by golden ratio to find next number
	phi := (1 + math.Sqrt(5)) / 2
	next := int(math.Round(float64(number) / phi))
	return next
}

func NextFibonacci(number int) int {
	// Fail if not fibonacci
	if !IsFibonacci(number) {
		return -1
	}
	// 0 is ambiguous, assuming unique number
	if number == 0 {
		return 1
	}
	// Multiply by golden ratio to find next number
	phi := (1 + math.Sqrt(5)) / 2
	next := int(math.Round(float64(number) * phi))
	return next
}
