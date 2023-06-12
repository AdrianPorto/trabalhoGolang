package main

import (
	"fmt"
	"sort"
	"strings"
)

// Crie um programa que receba duas strings informadas pelo usuário e imprima se elas são anagramas.

func main() {
	var str1, str2 string

	println("---------------------------------")
	println("| Digite a primeira string:     |")
	println("---------------------------------")
	fmt.Scanln(&str1)

	println("---------------------------------")
	println("| Digite a segunda string:      |")
	println("---------------------------------")
	fmt.Scanln(&str2)

	if saoAnagramas(str1, str2) {
		println("---------------------------------")
		println("| As strings são anagramas!     |")
		println("---------------------------------")
	} else {
		println("---------------------------------")
		println("| As strings não são anagramas! |")
		println("---------------------------------")
	}
}

func saoAnagramas(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	if len(str1) != len(str2) {
		return false
	}

	str1Sorted := sortString(str1)
	str2Sorted := sortString(str2)

	return str1Sorted == str2Sorted
}

func sortString(str string) string {
	strChars := strings.Split(str, "")

	sort.Strings(strChars)

	sortedStr := strings.Join(strChars, "")

	return sortedStr
}
