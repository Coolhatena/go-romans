package roman_numbers

import (
	"fmt"
)

func verifyCanBeConsequent(rm1 byte, rm2 byte) bool {
	valid_romans := map[byte]int{
		'M': 0,
		'D': 1,
		'C': 2,
		'L': 3,
		'X': 4,
		'V': 5,
		'I': 6,
	}

	return valid_romans[rm1] < valid_romans[rm2]
}

func verifySubstract(rm1 byte, rm2 byte) bool {
	if rm2 == 'V' || rm2 == 'X' {
		return rm1 == 'I'
	}

	if rm2 == 'L' || rm2 == 'C' {
		return rm1 == 'X'
	}

	if rm2 == 'M' {
		return rm1 == 'C'
	}

	return false
}

func verifyRomanSyntax(roman_num string) ([]int, bool) {
	actualChar := roman_num[0]
	consec_count := 0
	var break_indexes []int

	for i := 1; i < len(roman_num); i++ {
		if actualChar == roman_num[i] {
			consec_count++
			if consec_count > 2 {
				return []int{}, false
			}
		} else {
			consec_count = 0
			is_consec := verifyCanBeConsequent(actualChar, roman_num[i])

			if is_consec {
				break_indexes = append(break_indexes, i)
			} else {
				is_substract := verifySubstract(actualChar, roman_num[i])
				if !is_substract {
					return []int{}, false
				}
			}
		}

		actualChar = roman_num[i]
	}
	break_indexes = append(break_indexes, len(roman_num))
	return break_indexes, true
}

func sumRomans(roman_numbers []string) int {
	romans_table := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for _, roman := range roman_numbers {
		switch len(roman) {
		case 1:
			sum += romans_table[roman[0]]
		case 2:
			if verifySubstract(roman[0], roman[1]) {
				sum += romans_table[roman[1]] - romans_table[roman[0]]
			} else {
				sum += romans_table[roman[0]] + romans_table[roman[1]]
			}
		case 3:
			sum += romans_table[roman[0]] * 3
		}
	}

	return sum
}

func R2A(roman_num string) int {
	fmt.Println(roman_num)
	break_indexes, is_valid := verifyRomanSyntax(roman_num)

	if !is_valid {
		return -1
	}

	var individual_numbers []string
	start := 0
	for _, index := range break_indexes {
		individual_numbers = append(individual_numbers, roman_num[start:index])
		start = index
	}
	decimal_number := sumRomans(individual_numbers)
	return decimal_number
}
