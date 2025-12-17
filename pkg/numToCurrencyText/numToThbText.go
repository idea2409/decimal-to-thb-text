package numtocurrencytext

import (
	"strings"

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

// (recursive func)
func integerToThbText(num decimal.Decimal) string {
	// ----------- finish loop -----------
	// when no number to convert
	if num.Cmp(decimal.Zero) == 0 {
		return ""
	}

	// when going to be the last converting
	if num.Cmp(oneMillion) < 0 {
		return convertThbSixDigits(num)
	}
	// -------------------------------------

	// split to two parts -> low: last 6 digits, high: the prefix part
	high := num.Div(oneMillion).Floor()
	low := num.Mod(oneMillion)

	// recursive the prefix to have "ล้าน" at the end
	output := integerToThbText(high) + "ล้าน"

	// convert the last 6 digits
	if low.Cmp(decimal.Zero) > 0 {
		output += convertThbSixDigits(low)
	}

	return output
}

func convertThbSixDigits(num decimal.Decimal) string {
	var b strings.Builder
	mStr := num.String()
	mLength := len(mStr)

	for i, ch := range mStr {
		digit := byte(ch)
		unitIndex := mLength - i - 1

		if digit == '0' {
			continue
		}

		// handle special cases for Thai numbering rules
		if !(digit == '1' && unitIndex == 1) { // ignore the digit when "1" is at the "สิบ" position
			switch {
			case digit == '1' && unitIndex == 0 && mLength > 1: // handle when "1" is at the "เอ็ด" position
				b.WriteString("เอ็ด")
			case digit == '2' && unitIndex == 1: // handle when "2" is at the "ยี่" position
				b.WriteString("ยี่")
			default:
				b.WriteString(mapTextOfNum[string(digit)])
			}
		}
		// add unit
		b.WriteString(thaiUnits[unitIndex])
	}

	return b.String()
}
