package base62

import (
	"testing"
)

func areEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}

func Test_EncodeToString_With1000_ReturnsRS(t *testing.T) {
	num := 1000
	expected := "G8"

	actual := EncodeToString(num)

	areEqual(t, expected, actual)
}

func Test_EncodeToString_With0_Returns0(t *testing.T) {
	num := 0
	expected := "0"

	actual := EncodeToString(num)

	areEqual(t, expected, actual)
}

func Test_EncodeToString_With10_ReturnsA(t *testing.T) {
	num := 10
	expected := "A"

	actual := EncodeToString(num)

	areEqual(t, expected, actual)
}

func Test_EncodeToString_With35_ReturnsZ(t *testing.T) {
	num := 35
	expected := "Z"

	actual := EncodeToString(num)

	areEqual(t, expected, actual)
}

func Test_EncodeToString_With36_Returns10(t *testing.T) {
	num := 36
	expected := "a"

	actual := EncodeToString(num)

	areEqual(t, expected, actual)
}

func Test_EncodeToString_With64_Returns1S(t *testing.T) {
	num := 64
	expected := "12"

	actual := EncodeToString(num)

	areEqual(t, expected, actual)
}

func Test_DecodeToInt_WithRS_Returns1000(t *testing.T) {
	base36 := "G8"
	expected := 1000

	actual := DecodeToInt(base36)

	areEqual(t, expected, actual)
}

func Test_DecodeToInt_With1S_Returns64(t *testing.T) {
	base36 := "12"
	expected := 64

	actual := DecodeToInt(base36)

	areEqual(t, expected, actual)
}
