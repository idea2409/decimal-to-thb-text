package main

import (
	numtocurrencytext "decimal-to-thb-text/pkg/numToCurrencyText"
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// --------- added lines ---------
	veryLargeNum, err := decimal.NewFromString("123456123456123456.99")
	if err != nil {
		panic(err)
	}
	// -------------------------------

	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
		veryLargeNum, // added lines
	}

	for _, input := range inputs {
		fmt.Println(input.String())

		// convert decimal to thai text (baht) and print the result here

		fmt.Println(numtocurrencytext.NumToThbText(input))
		fmt.Println("--------------------------------")
	}
}
