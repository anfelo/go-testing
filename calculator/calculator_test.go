package calculator_test

import (
	"testing"

	"github.com/anfelo/go-testing/calculator"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	// Sub tests
	t.Run("test for all 3 digit armstrong numbers", func(t *testing.T) {
		// Test tables
		tests := []TestCase{
			{name: "Testing value for: 153", value: 153, expected: true},
			{name: "Testing value for: 370", value: 370, expected: true},
			{name: "Testing value for: 371", value: 371, expected: true},
			{name: "Testing value for: 407", value: 407, expected: true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := calculator.CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fatalf("Expected %t but got %t", test.expected, actual)
				}
			})
		}
	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("should fail for 350", func(t *testing.T) {
		tc := TestCase{
			value:    350,
			expected: false,
		}

		tc.actual = calculator.CalculateIsArmstrong(tc.value)
		if tc.actual != tc.expected {
			t.Fatalf("Expected %t but got %t", tc.expected, tc.actual)
		}
	})

	t.Run("should fail for 300", func(t *testing.T) {
		tc := TestCase{
			value:    300,
			expected: false,
		}

		tc.actual = calculator.CalculateIsArmstrong(tc.value)
		if tc.actual != tc.expected {
			t.Fatalf("Expected %t but got %t", tc.expected, tc.actual)
		}
	})
}

func benchmarkCalculateIsArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.CalculateIsArmstrong(input)
	}
}

func BenchmarkCalculateIsArmstrong370(b *testing.B) {
	benchmarkCalculateIsArmstrong(370, b)
}

func BenchmarkCalculateIsArmstrong371(b *testing.B) {
	benchmarkCalculateIsArmstrong(371, b)
}

func BenchmarkCalculateIsArmstrong0(b *testing.B) {
	benchmarkCalculateIsArmstrong(0, b)
}
