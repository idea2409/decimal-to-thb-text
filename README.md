# decimal-to-thb-text

This project converts numeric values (`decimal.Decimal`) into **Thai Baht text format** (บาท / สตางค์).

It supports:

- Arbitrary-precision numbers using `github.com/shopspring/decimal`
- Proper Thai numbering rules (เอ็ด, ยี่, สิบ, ร้อย, พัน, หมื่น, ...)
- Automatic separation of integer (บาท) and fractional (สตางค์) parts
- Rounding down fractional values to 2 decimal places
- Negative numbers

---

## ✨ Features

- Convert `decimal.Decimal` to Thai currency text
- Correct handling of Thai numeric special grammar:
  - "1" at the "เอ็ด" position
  - "1" at the 'ten' position
  - "2" at the "ยี่" position
  - the suffix "ถ้วน" when no fractional part
- Negative numbers will be tagged as "ติดลบ"
- Zero will be returned as "ศูนย์บาทถ้วน"
- Simple API: `NumToThbText(decimal.Decimal)`
- Supports values from **0.00** up to very large numbers (works as long as the input is decimal.Decimal)

---

## 🧠 Example

```go
package main

import (
	numtocurrencytext "decimal-to-thb-text/pkg/numToCurrencyText"
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	input1 := decimal.RequireFromString("1234")
	input2 := decimal.RequireFromString("621908501.25")
	input3, _ := decimal.NewFromString("123456123456123456.99")
	fmt.Println(numtocurrencytext.NumToThbText(input1))
	fmt.Println(numtocurrencytext.NumToThbText(input2))
	fmt.Println(numtocurrencytext.NumToThbText(input3))
}
```

Output:

```
หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน
หกร้อยยี่สิบเอ็ดล้านเก้าแสนแปดพันห้าร้อยเอ็ดบาทยี่สิบห้าสตางค์
หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านหนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านหนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกบาทเก้าสิบเก้าสตางค์
```

---

## 🚀 How to Run

Make sure you have **Go** installed (Go 1.20+ recommended).

Run the application using:

```bash
go run main/main.go
```

## 🚀 How to Run Tests

Run tests for the numtocurrencytext package:

```bash
go test ./pkg/numToCurrencyText
```

## 📂 Project Structure

.
├── go.mod
├── go.sum
├── main/
│ └── main.go
├── pkg/
│ └── numToCurrencyText/
│ ├── converter.go
│ └── converter_test.go
└── README.md
