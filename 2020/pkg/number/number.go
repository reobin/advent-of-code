package number

import (
	"strconv"
)

// Range represents an integer range from min to max
type Range struct {
	Minimum int
	Maximum int
}

// ComputeProduct takes in a list of numbers and computes their product
func ComputeProduct(numbers []int) int {
	product := 1
	for _, n := range numbers {
		product *= n
	}
	return product
}

// Contains returns true if the number was found in the list
func Contains(numbers []int, element int) bool {
	for _, n := range numbers {
		if element == n {
			return true
		}
	}
	return false
}

// ConvertToNumbers converts a list of strings to a list of numbers
// Note: It ignores the values that cannot be converted
func ConvertToNumbers(numbersStr []string) []int {
	var numbers []int
	for _, value := range numbersStr {
		n, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		numbers = append(numbers, n)
	}
	return numbers
}

// IsNumberInRange returns true if the number is within the range
func IsNumberInRange(value int, valueRange Range) bool {
	return value >= valueRange.Minimum && value <= valueRange.Maximum
}
