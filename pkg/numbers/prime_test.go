package numbers_test

import (
	"testing"

	"github.com/Revanee/fibonacci-api/pkg/numbers"
)

func TestIsPrime(t *testing.T) {

	firstPrimeNumbers := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	}
	for _, primeNumber := range firstPrimeNumbers {
		if !numbers.IsPrime(primeNumber) {
			t.Errorf("%d should be recognized as prime", primeNumber)
		}
	}

	firstNonPrimeNumbers := []int{
		4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20,
	}
	for _, nonPrimeNumber := range firstNonPrimeNumbers {
		if numbers.IsPrime(nonPrimeNumber) {
			t.Errorf("%d should not be recognized as prime", nonPrimeNumber)
		}
	}
}
