package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
	numbers           = "1234567890"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns an integer between the min, max range inclusive
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //nolint
}

// RandomStringValid returns a string of length n consisting a combination of alphanumeric characters
// that is semantically valid for the /username endpoint and should pass the validation scheme
func RandomStringValid(n int) string {
	var sb strings.Builder
	k := len(lowerCaseAlphabet)
	l := len(numbers)

	for i := 0; i < n; i++ {
		// for the first character choose from either upper or lower case alphabets
		if i == 0 {
			c := rand.Intn(2) //nolint
			if c == 0 {
				sb.WriteByte(upperCaseAlphabet[rand.Intn(k)]) //nolint
				continue
			}
			sb.WriteByte(lowerCaseAlphabet[rand.Intn(k)]) //nolint
		}

		c := rand.Intn(3) //nolint
		switch c {
		case 0:
			sb.WriteByte(upperCaseAlphabet[rand.Intn(k)]) //nolint
		case 1:
			sb.WriteByte(lowerCaseAlphabet[rand.Intn(k)]) //nolint
		case 2:
			sb.WriteByte(numbers[rand.Intn(l)]) //nolint
		}
	}

	return sb.String()
}

// RandomStringInvalid returns a string of length n consisting a combination of alphanumeric characters
// that is semantically invalid for the /username endpoint and should not pass validation scheme
func RandomStringInvalid(n int) string {
	var sb strings.Builder
	k := len(lowerCaseAlphabet)
	l := len(numbers)

	for i := 0; i < n; i++ {
		// if length of the string is between 8 to 12 characters, let the first character be a number
		// so that it is invalid for the API endpoint
		if n >= 8 && n <= 12 && i == 0 {
			sb.WriteByte(numbers[rand.Intn(l)]) //nolint
			continue
		}

		c := rand.Intn(3) //nolint
		switch c {
		case 0:
			sb.WriteByte(upperCaseAlphabet[rand.Intn(k)]) //nolint
		case 1:
			sb.WriteByte(lowerCaseAlphabet[rand.Intn(k)]) //nolint
		case 2:
			sb.WriteByte(numbers[rand.Intn(l)]) //nolint
		}
	}

	return sb.String()
}
