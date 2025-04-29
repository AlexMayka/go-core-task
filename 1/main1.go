package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func printTypeOfValues(values ...any) {
	for _, value := range values {
		fmt.Printf("Значение: %v, тип: %T\n", value, value)
	}
}

func combiningValuesToStr(values ...any) string {
	var builder strings.Builder
	for _, value := range values {
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	return builder.String()
}

func convertToRunes(input string) []rune {
	return []rune(input)
}

func hashRunesSHA256(runes []rune) string {
	salt := []rune("go-2024")
	mid := len(runes) / 2
	runesWithSole := append(runes[:mid], append(salt, runes[mid:]...)...)

	hash := sha256.New()
	hash.Write([]byte(string(runesWithSole)))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	var numDecimal int = 42       // Десятичная система
	var numOctal int = 052        // Восьмеричная система
	var numHexadecimal int = 0x2A // Шестнадцатиричная система

	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	printTypeOfValues(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	unionStr := combiningValuesToStr(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	runes := convertToRunes(unionStr)
	hash := hashRunesSHA256(runes)

	fmt.Println(hash)
}
