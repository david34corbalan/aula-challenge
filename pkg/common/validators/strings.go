package validators

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func Strings(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	for _, char := range str {
		if !(unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsPunct(char) || unicode.IsSpace(char) || isExtendedASCII(char)) {
			return false
		}
	}
	return true
}

// isExtendedASCII checks if a character is an extended ASCII character (128-255)
func isExtendedASCII(char rune) bool {
	return char >= 128 && char <= 255
}
