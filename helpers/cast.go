package helpers

import (
	"fmt"
	"strconv"
)

func Atoi(stringValue string) int {
	integerValue, err := strconv.Atoi(stringValue)
	if err != nil {
		// ... handle error
		errorMessage := fmt.Sprintf("The conversio of %s to int, resulted in this error %s", stringValue, err)
		panic(errorMessage)
	}
	return integerValue
}

func Rtoi(runeValue rune) int {
	return Atoi(string(runeValue))
}

func Btoi(byteValue byte) int {
	return Atoi(string(byteValue))
}

func Atoiv(stringValue string) []int {
	var returnValue = []int{}
	for _, element := range stringValue {
		returnValue = append(returnValue, Rtoi(element))
	}
	return returnValue
}

func Ivtosv(intValues []int) []string {
	var returnValue = []string{}
	for _, element := range intValues {
		returnValue = append(returnValue, fmt.Sprint(element))
	}
	return returnValue
}
