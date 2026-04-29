package affinecipher

import (
	"errors"
	"unicode"
)

// M is a special variable in the cipher
var M int = 26

func Encode(text string, a, b int) (string, error) {
	if gcd(a, M) != 1 {
		return "", errors.New("a and b must be coprime")
	}
	var result []rune
	for _, ch := range text {
		if unicode.IsLetter(ch) {
			x := int(unicode.ToLower(ch)) - 'a'
			encoded := (a*x + b) % M
			result = append(result, rune(encoded+'a'))
		} else if unicode.IsDigit(ch) {
			result = append(result, ch)
		}
	}
	return group(result, 5), nil
}

func Decode(text string, a, b int) (string, error) {
	if gcd(a, M) != 1 {
		return "", errors.New("a and b must be coprime")
	}

	aInv := mmi(a, 26)
	var result []rune
	for _, ch := range text {
		if unicode.IsLetter(ch) {
			y := int(ch) - 'a'
			decoded := (aInv * (y - b + M*M)) % M
			result = append(result, rune(decoded+'a'))
		} else if unicode.IsDigit(ch) {
			result = append(result, ch)
		}
	}

	return string(result), nil
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmi(a, m int) int {
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return 1
}

func group(parts []rune, spacing int) string {
	final := ""
	for i, ch := range parts {
		final += string(ch)
		if (i+1)%spacing == 0 && i < len(parts)-1 {
			final += " "
		}
	}
	return final
}
