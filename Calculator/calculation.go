package main

import "strconv"

func arabic_calc(x, y, operator string) int {
	int_x, err_x := strconv.Atoi(x)
	if err_x != nil {
		panic(err_x)
	}

	int_y, err_y := strconv.Atoi(y)
	if err_y != nil {
		panic(err_y)
	}

	if operator == "+" {
		return int_x + int_y
	} else if operator == "-" {
		return int_x - int_y
	} else if operator == "*" {
		return int_x * int_y
	} else if operator == "/" {
		return int_x / int_y
	} else {
		panic("Ошибка! Неизвестный оператор")
	}
}

func roman_calc(x, y, operator string) string {
	var rom_unit = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	var rom_decade = map[int]string{
		0: "", 1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC", 10: "C",
	}

	num1, num2 := rom_unit[x], rom_unit[y]
	result := arabic_calc(strconv.Itoa(num1), strconv.Itoa(num2), operator)
	if result <= 0 {
		panic("Ошибка! Результат вычисления нельзя представить в римской системе счисления")
	}
	var total = ""
	dec, unit := result/10, result%10
	total += rom_decade[dec]
	for k, v := range rom_unit {
		if v == unit {
			total += k
		}
	}
	return total
}
