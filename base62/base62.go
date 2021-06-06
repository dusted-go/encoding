package base62

import (
	"math"
	"unicode"
)

// EncodeToString encodes an integer into a base62 string.
// Base62 is a case-sensitive string consisting of 0-9, A-Z and a-z characters.
func EncodeToString(num int) string {
	var convertFunc func(int, string) string

	convertFunc = func(num int, str string) string {
		base := 62
		q := num / base
		r := num % base

		// 0-9
		dec := r + 48

		// A-Z
		if r >= 10 {
			dec = r + 55
		}

		// a-z
		if r >= 10+26 {
			dec = r + 61
		}

		str = string(rune(dec)) + str

		if q == 0 {
			return str
		}

		return convertFunc(q, str)
	}

	return convertFunc(num, "")
}

// DecodeToInt decodes a base62 string into the original integer.
func DecodeToInt(str string) int {
	runes := []rune(str)
	result := 0
	var base float64 = 62
	var exponent float64 = 0

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		d := int(r)

		switch {
		case unicode.IsDigit(r):
			d -= 48
		case unicode.IsLetter(r) && unicode.IsUpper(r):
			d -= 55
		case unicode.IsLetter(r) && unicode.IsLower(r):
			d -= 61
		}
		result += d * int(math.Pow(base, exponent))
		exponent++
	}

	return result
}
