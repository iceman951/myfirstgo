package Calculator

import (
	"math"
)

type CalculatorInterface interface {
	// GetCalculatorName() string
	// Add(x, y int) int
	// Subtract(x, y int) int
	// Multiply(x, y int) int
	// Divide(x, y int) int

	Power(x, y int) float64
	Square(x int) float64
	IsOdd(x int) bool
}

type MyCalculator struct {
	name string
}

func NewTheBestCalculator(name string) CalculatorInterface {
	return &MyCalculator{
		name: name,
	}
}

func (m *MyCalculator) GetCalculatorName() string {
	return m.name
}

func (m MyCalculator) Add(x, y int) int {
	return x + y
}

func (m MyCalculator) Subtract(x, y int) int {
	return x - y
}

func (m MyCalculator) Multiply(x, y int) int {
	return x * y
}

func (m MyCalculator) Divide(x, y int) int {
	return x / y
}

func (m MyCalculator) Square(input int) float64 {
	return math.Sqrt(float64(input))
}

func (m MyCalculator) Power(x, y int) float64 {
	return math.Pow(float64(x), float64(y))
}

func (m MyCalculator) IsOdd(x int) bool {
	if x%2 == 0 {
		return false
	} else {
		return true
	}
}
