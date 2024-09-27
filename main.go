package main

import (
	"fmt"
	"calc"
)

func main() {

	var num1 float64
	var err error
	fmt.Println("Введите первое число:")
	_, err = fmt.Scanln(&num1)
	if err != nil {
		fmt.Printf("ошибка при сканировании первого числа: %v", err)
	}

	var operator string
	fmt.Println("Введите оператор:")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Printf("ошибка при сканировании оператора: %v", err)
	}

	var num2 float64
	fmt.Println("Введите второе число:")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Printf("ошибка при сканировании второго числа: %v", err)
	}

	calculator := calc.NewCalculator()
	result := calculator.Calculate(num1, num2, operator)

	fmt.Println(result)
}