package main

import (
	"fmt"
	"romans/roman_numbers"
)

func main() {
	fmt.Println("Roman to Arabic:")
	roman := "MMMCMXCIX"
	r2a := roman_numbers.R2A(roman)
	fmt.Println(r2a)
	fmt.Println("Arabic to Roman:")
	a2r := roman_numbers.A2R(r2a)
	fmt.Println(a2r)

}
