package numtocurrencytext

import (
	"fmt"
	"slices"

	"github.com/shopspring/decimal"
)

func NumToThbText(num decimal.Decimal) string {
	// split to two parts (integer & fractional)
	fractional := num.Mod(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100)).Truncate(0) // round down when the fractional part has more than two digits
	integer := num.Truncate(0)

	// return when num is 0
	if integer.Cmp(decimal.Zero) == 0 && fractional.Cmp(decimal.Zero) == 0 {
		return "ศูนย์บาทถ้วน"
	}

	// init output
	output := ""

	// handle negative number
	if integer.Cmp(decimal.Zero) == -1 {
		output += "ติดลบ"
		integer = integer.Neg()
		fractional = fractional.Neg()
	}

	// convert the integer part
	if integer.Cmp(decimal.Zero) == 1 {
		output += integerToThbText(integer) + "บาท"
	}

	// convert the fractional part
	if fractional.Cmp(decimal.Zero) == 1 { // If the value has a fractional part, convert the fractional part into Thai text representing "สตางค์"
		output += integerToThbText(fractional) + "สตางค์"
	} else { // If the value has no fractional part, append the suffix "ถ้วน" to the result.
		output += "ถ้วน"
	}

	return output
}

var (
	oneMillion   = decimal.NewFromInt(1000000)
	thaiUnits    = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}
	mapTextOfNum = map[string]string{
		"1": "หนึ่ง",
		"2": "สอง",
		"3": "สาม",
		"4": "สี่",
		"5": "ห้า",
		"6": "หก",
		"7": "เจ็ด",
		"8": "แปด",
		"9": "เก้า",
	}
)

func integerToThbText(integer decimal.Decimal) string {
	output := ""

	// split to array of 6-digit group
	millions := []decimal.Decimal{}
	for integer.Cmp(oneMillion) != -1 {
		millions = append(millions, integer.Mod(oneMillion))
		integer = integer.Div(oneMillion).Floor()
	}
	millions = append(millions, integer)

	// reverse array to make it be sorted
	slices.Reverse(millions)

	for i, m := range millions {
		for j, v := range m.String() {
			digit := fmt.Sprintf("%c", v)
			unitIndex := len(m.String()) - j - 1

			// skip reading digit 0
			if digit == "0" {
				continue
			}

			// handle special cases for Thai numbering rules
			if digit == "1" && unitIndex == 0 && len(m.String()) > 1 { // handle when "1" is at the "เอ็ด" position
				output += "เอ็ด"
			} else if digit == "2" && unitIndex == 1 { // handle when "2" is at the "ยี่" position
				output += "ยี่"
			} else if !(digit == "1" && unitIndex == 1) { // ignore the digit when "1" is at the "สิบ" position
				output += mapTextOfNum[digit]
			}

			// add unit
			output += thaiUnits[unitIndex]
		}

		// add when not the last million group
		if i < len(millions)-1 {
			output += "ล้าน"
		}
	}

	return output
}
