package propertybasedtests

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToArabic(roman string) uint16 {
	var res uint16 = 0

	for len(roman) > 0 {
		for _, numeral := range allRomanNumerals {
			if strings.HasPrefix(roman, numeral.Symbol) {
				res += uint16(numeral.Value)
				roman = strings.TrimPrefix(roman, numeral.Symbol)
			}
		}
	}

	return res
}

func ConvertToRoman(arabic uint16) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
