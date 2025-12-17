package main

import (
	"fmt"
	"slices"

	"github.com/shopspring/decimal"
)

// Develop a function to convert a decimal value to Thai text with a "baht" currency suffix.
// Please take into consideration how this code should be integrated into a service.

// The input will be a decimal value using the github.com/shopspring/decimal package.
// The output must be a string in Thai text format.
// If the value has no fractional part, append the suffix "ถ้วน" to the result.
// If the value has a fractional part, convert the fractional part into Thai text representing "สตางค์"

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(0.00), // min
		decimal.NewFromFloat(21654321654321.12),
		decimal.NewFromFloat(99999999999999.98), // max
	}

	for _, input := range inputs {
		fmt.Println(input.String())
		// convert decimal to thai text (baht) and print the result here
		fmt.Println(floatToThbText(input))

		fmt.Println("--------------------------------")
	}
}

func floatToThbText(num decimal.Decimal) string {
	output := ""
	fractional := num.Mod(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
	integer := num.Truncate(0)
	fmt.Println("integer = ", integer)
	fmt.Println("fractional = ", fractional)

	if integer.Cmp(decimal.Zero) == 1 {
		output += integerToThbText(integer) + "บาท"
	} else {
		output = "ศูนย์บาทถ้วน"
	}

	if fractional.Cmp(decimal.Zero) == 1 { // If the value has a fractional part, convert the fractional part into Thai text representing "สตางค์"
		output += integerToThbText(fractional) + "สตางค์"
	} else { // If the value has no fractional part, append the suffix "ถ้วน" to the result.
		output += "ถ้วน"
	}

	return output
}

var (
	mapTextOfNum = map[string]string{
		// "0": "",
		"1": "หนึี่ง",
		"2": "สอง",
		"3": "สาม",
		"4": "สี่",
		"5": "ห้า",
		"6": "หก",
		"7": "เจ็ด",
		"8": "แปด",
		"9": "เก้า",
	}
	nuais = []string{
		"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน",
	}
)

func integerToThbText(integer decimal.Decimal) string {
	output := ""

	// split to millions
	millions := []decimal.Decimal{}
	for integer.Cmp(decimal.NewFromInt(999999)) == 1 {
		millions = append(millions, integer.Mod(decimal.NewFromInt(1000000)))
		integer = integer.Div(decimal.NewFromInt(1000000)).Floor()
	}
	millions = append(millions, integer)

	slices.Reverse(millions)

	for i, m := range millions {
		// fmt.Println(m)

		for j, v := range m.String() {
			digit := fmt.Sprintf("%c", v)
			nuaiIndex := len(m.String()) - j - 1

			// skip reading 0
			if digit == "0" {
				continue
			}

			// todo:
			// 1 "เอ็ด(หน่วย)" หรือ "หนึ่ง(หน่วย)" , สิบเฉยๆ
			// 2 "ยี่(สิบ)" หรือ "สอง(หน่วย)"
			// ...
			if digit == "1" && nuaiIndex == 0 && len(m.String()) > 1 {
				output += "เอ็ด"
			} else if digit == "1" && nuaiIndex == 1 && len(m.String()) > 1 {
				// no adding
			} else if digit == "2" && nuaiIndex == 1 {
				output += "ยี่"
			} else {
				output += mapTextOfNum[digit]
			}

			output += nuais[nuaiIndex]
		}

		if i < len(millions)-1 {
			output += "ล้าน"
		}
	}

	return output
}
