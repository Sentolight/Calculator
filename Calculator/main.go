package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию:")
	input, _ := reader.ReadString('\n')
	elements := strings.Fields(input)

	if len(elements) != 3 {
		panic("Ошибка! Ожидается формат: 'число' 'оператор' 'число'")
	}

	num1, operator, num2 := elements[0], elements[1], elements[2]

	if (is_arabic(num1)+is_arabic(num2) != 2) && (is_roman(num1)+is_roman(num2) != 2) {
		panic("Ошибка! Допустимые числа могут быть только римскими или только арабскими в диапазоне от 1 до 10")
	}

	if is_operator(operator) != 1 {
		panic("Ошибка! Допустимые операторы: '+', '-', '*', '/'")
	}

	if is_arabic(num1)+is_arabic(num2) == 2 {
		fmt.Println(arabic_calc(num1, num2, operator))
	}

	if is_roman(num1)+is_roman(num2) == 2 {
		fmt.Println(roman_calc(num1, num2, operator))
	}
}
