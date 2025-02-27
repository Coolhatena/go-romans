package roman_numbers

import (
	"strconv"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func romanCalc(num int, min string, mid string, max string) string {
	if num < 4 {
		return strings.Repeat(min, num)
	}

	if num == 4 {
		return min + mid
	}

	if num < 9 {
		return mid + strings.Repeat(min, num-5)
	} else {
		return min + max
	}
}

func A2R(num int) string {
	roman_num := ""
	string_num := reverseString(strconv.Itoa(num))
	characters_quant := len(string_num)
	if characters_quant > 4 {
		return "Only numbers below 3999 can be converted"
	}

	count := 0

	for count < characters_quant {
		if count == 0 {
			roman_num = romanCalc(int(string_num[count]-'0'), "I", "V", "X")
		}

		if count == 1 {
			roman_num = romanCalc(int(string_num[count]-'0'), "X", "L", "C") + roman_num
		}

		if count == 2 {
			roman_num = romanCalc(int(string_num[count]-'0'), "C", "D", "M") + roman_num
		}

		if count == 3 {
			roman_num = romanCalc(int(string_num[count]-'0'), "M", "V", "X") + roman_num
		}

		count++
	}

	return roman_num
}
