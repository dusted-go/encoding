package base36

import (
	"math"
	"strings"
	"unicode"
)

// EncodeToString encodes an integer into a base36 string.
// Base36 is a case-insensitive string consisting of 0-9 and A-Z characters.
func EncodeToString(num int) string {
	var convertFunc func(int, string) string

	convertFunc = func(num int, str string) string {
		q := num / 36
		r := num % 36

		// 0-9
		dec := r + 48

		// A-Z
		if r >= 10 {
			dec = r + 55
		}

		str = string(rune(dec)) + str

		if q == 0 {
			return str
		}

		return convertFunc(q, str)
	}

	return convertFunc(num, "")
}

// DecodeToInt decodes a base36 string into the original integer.
func DecodeToInt(str string) int {
	runes := []rune(strings.ToUpper(str))
	result := 0
	var numBase float64 = 36
	var exponent float64 = 0

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		d := int(r)

		if unicode.IsDigit(r) {
			d -= 48
		} else {
			d -= 55
		}
		result += d * int(math.Pow(numBase, exponent))
		exponent++
	}

	return result
}
