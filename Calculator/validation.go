package main

func is_arabic(num string) int {
	num_list := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for i := 0; i < len(num_list); i++ {
		valid_num := num_list[i]
		if valid_num == num {
			return 1
		}
	}
	return 0
}

func is_roman(num string) int {
	num_list := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(num_list); i++ {
		valid_num := num_list[i]
		if valid_num == num {
			return 1
		}
	}
	return 0
}

func is_operator(num string) int {
	num_list := [4]string{"+", "-", "*", "/"}
	for i := 0; i < len(num_list); i++ {
		valid_num := num_list[i]
		if valid_num == num {
			return 1
		}
	}
	return 0
}
