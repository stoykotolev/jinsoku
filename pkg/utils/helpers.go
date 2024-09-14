package utils

import "math/rand"

var SpecialSymbolsAndNumbers = []rune{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	'!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
	'-', '=', '[', ']', '\\', ';', '\'', ',', '.', '/',
	'_', '+', '{', '}', '|', ':', '"', '<', '>', '?',
	'`', '~',
}

func randomInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func GetRandomSymbol() string {
	arrLen := len(SpecialSymbolsAndNumbers)
	el := randomInRange(0, arrLen-1)

	// The character to match
	return string(SpecialSymbolsAndNumbers[el])
}
