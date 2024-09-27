package calc

import "log"

const (
	operatorAddition       = "+"
	operatorSubtraction    = "-"
	operatorMultiplication = "*"
	operatorDivision       = "/"
)

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

func (c *calculator) Calculate(num1, num2 float64, operator string) float64 {
	switch operator {
	case operatorAddition:
		return c.add(num1, num2)
	case operatorSubtraction:
		return c.subtract(num1, num2)
	case operatorMultiplication:
		return c.multiply(num1, num2)
	case operatorDivision:
		return c.divide(num1, num2)
	default:
		log.Printf("несуществующий оператор: %s\n", operator)
		return 0
	}
}

func (c *calculator) add(num1, num2 float64) float64 {
	return num1 + num2
}
func (c *calculator) subtract(num1, num2 float64) float64 {
	return num1 - num2
}
func (c *calculator) multiply(num1, num2 float64) float64 {
	return num1 * num2
}
func (c *calculator) divide(num1, num2 float64) float64 {
	if num2 == 0 {
		log.Println("ошибка, деление на 0")
		return 0
	}

	return num1 / num2
}