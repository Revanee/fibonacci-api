package numbers_test

import (
	"testing"

	"github.com/Revanee/fibonacci-api/pkg/numbers"
)

func TestIsFibonacci(t *testing.T) {

	firstFibonacciNumbers := []int{
		0, 1, 1, 2, 3, 5, 8, 13, 21, 34,
	}
	for _, fibonacciNumber := range firstFibonacciNumbers {
		if !numbers.IsFibonacci(fibonacciNumber) {
			t.Errorf("%d should be recognized as Fibonacci", fibonacciNumber)
		}
	}

	firstNonFibonacciNumbers := []int{
		4, 6, 7, 9, 10, 11, 12, 14, 15, 16,
	}
	for _, nonFibonacciNumber := range firstNonFibonacciNumbers {
		if numbers.IsFibonacci(nonFibonacciNumber) {
			t.Errorf("%d should not be recognized as Fibonacci", nonFibonacciNumber)
		}
	}

	// False for negative numbers
	negativeFibonacciNumbers := []int{
		-8, -3, -1, -10,
	}
	for _, negativeFibonacci := range negativeFibonacciNumbers {
		if numbers.IsFibonacci(negativeFibonacci) {
			t.Errorf("%d should not be recognized as Fibonacci", negativeFibonacci)
		}
	}
}

func TestIsPerfectSquare(t *testing.T) {
	firstPerfectSquares := []int{
		1, 4, 9, 16, 25, 36, 49, 64, 81, 100,
	}
	for _, perfectSquare := range firstPerfectSquares {
		if !numbers.IsPerfectSquare(perfectSquare) {
			t.Errorf("%d should be recognized as perfect square", perfectSquare)
		}
	}

	firstNonPerfectSquares := []int{
		2, 3, 5, 6, 7, 8, 10, 11, 12, 13, -2, -9,
	}
	for _, nonPerfectSquare := range firstNonPerfectSquares {
		if numbers.IsPerfectSquare(nonPerfectSquare) {
			t.Errorf("%d should not be recognized as perfect square", nonPerfectSquare)
		}
	}
}

func TestNextFibonacci(t *testing.T) {
	// 1 is ambiguous
	firstFibonacciNumbers := []int{
		0, 1, 2, 3, 5, 8, 13, 21, 34,
	}
	for i := 0; i < len(firstFibonacciNumbers)-1; i++ {
		current := firstFibonacciNumbers[i]
		next := numbers.NextFibonacci(current)
		expectedNext := firstFibonacciNumbers[i+1]
		if next != expectedNext {
			t.Errorf("The next Fibonacci after %d is %d, instead got %d", current, expectedNext, next)
		}
	}
}

func TestPreviousFibonacci(t *testing.T) {
	// 1 is ambiguous
	firstFibonacciNumbers := []int{
		0, 1, 2, 3, 5, 8, 13, 21, 34,
	}
	for i := 1; i < len(firstFibonacciNumbers); i++ {
		current := firstFibonacciNumbers[i]
		previous := numbers.PreviousFibonacci(current)
		expectedPrevious := firstFibonacciNumbers[i-1]
		if previous != expectedPrevious {
			t.Errorf("The previous Fibonacci before %d is %d, instead got %d", current, expectedPrevious, previous)
		}
	}
}
