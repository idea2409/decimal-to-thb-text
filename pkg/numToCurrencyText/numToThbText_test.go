package numtocurrencytext

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNumToThbText(t *testing.T) {
	tests := []struct {
		name     string
		input    decimal.Decimal
		expected string
	}{
		{
			name:     "zero",
			input:    decimal.NewFromInt(0),
			expected: "ศูนย์บาทถ้วน",
		},
		{
			name:     "single digit",
			input:    decimal.NewFromInt(1),
			expected: "หนึ่งบาทถ้วน",
		},
		{
			name:     "ten",
			input:    decimal.NewFromInt(10),
			expected: "สิบบาทถ้วน",
		},
		{
			name:     "eleven",
			input:    decimal.NewFromInt(11),
			expected: "สิบเอ็ดบาทถ้วน",
		},
		{
			name:     "twenty",
			input:    decimal.NewFromInt(20),
			expected: "ยี่สิบบาทถ้วน",
		},
		{
			name:     "twenty one",
			input:    decimal.NewFromInt(21),
			expected: "ยี่สิบเอ็ดบาทถ้วน",
		},
		{
			name:     "one hundred",
			input:    decimal.NewFromInt(100),
			expected: "หนึ่งร้อยบาทถ้วน",
		},
		{
			name:     "one million",
			input:    decimal.NewFromInt(1000000),
			expected: "หนึ่งล้านบาทถ้วน",
		},
		{
			name:     "fractional only",
			input:    decimal.NewFromFloat(0.91),
			expected: "เก้าสิบเอ็ดสตางค์",
		},
		{
			name:     "no fractional",
			input:    decimal.NewFromFloat(1234),
			expected: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน",
		},
		{
			name:     "with fractional",
			input:    decimal.NewFromFloat(33333.75),
			expected: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์",
		},
		{
			name:     "include 0 in the middle",
			input:    decimal.NewFromFloat(12034.75),
			expected: "หนึ่งหมื่นสองพันสามสิบสี่บาทเจ็ดสิบห้าสตางค์",
		},
		{
			name:     "include 0 at the end of integer part",
			input:    decimal.NewFromFloat(12030.75),
			expected: "หนึ่งหมื่นสองพันสามสิบบาทเจ็ดสิบห้าสตางค์",
		},
		{
			name:     "include 0 at the end of fractional part",
			input:    decimal.NewFromFloat(12034.70),
			expected: "หนึ่งหมื่นสองพันสามสิบสี่บาทเจ็ดสิบสตางค์",
		},
		{
			name:     "include 0 at the end of both integer and fractional part",
			input:    decimal.NewFromFloat(12030.70),
			expected: "หนึ่งหมื่นสองพันสามสิบบาทเจ็ดสิบสตางค์",
		},
		{
			name:     "include 1 at the end of integer part",
			input:    decimal.NewFromFloat(54331.75),
			expected: "ห้าหมื่นสี่พันสามร้อยสามสิบเอ็ดบาทเจ็ดสิบห้าสตางค์",
		},
		{
			name:     "include 1 at the end of fractional part",
			input:    decimal.NewFromFloat(54330.71),
			expected: "ห้าหมื่นสี่พันสามร้อยสามสิบบาทเจ็ดสิบเอ็ดสตางค์",
		},
		{
			name:     "include 1 at the end of both integer and fractional part",
			input:    decimal.NewFromFloat(54331.71),
			expected: "ห้าหมื่นสี่พันสามร้อยสามสิบเอ็ดบาทเจ็ดสิบเอ็ดสตางค์",
		},
		{
			name:     "complicated number",
			input:    decimal.NewFromFloat(621908501.25),
			expected: "หกร้อยยี่สิบเอ็ดล้านเก้าแสนแปดพันห้าร้อยเอ็ดบาทยี่สิบห้าสตางค์",
		},
		{
			name:     "negative number",
			input:    decimal.NewFromFloat(-621908501.25),
			expected: "ติดลบหกร้อยยี่สิบเอ็ดล้านเก้าแสนแปดพันห้าร้อยเอ็ดบาทยี่สิบห้าสตางค์",
		},
		{
			name:     "large number",
			input:    decimal.NewFromFloat(50780921908501.25),
			expected: "ห้าสิบล้านเจ็ดแสนแปดหมื่นเก้าร้อยยี่สิบเอ็ดล้านเก้าแสนแปดพันห้าร้อยเอ็ดบาทยี่สิบห้าสตางค์",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumToThbText(tt.input)
			if result != tt.expected {
				t.Errorf(
					"NumToThbText(%d) = %q, want %q",
					tt.input,
					result,
					tt.expected,
				)
			}
		})
	}
}
