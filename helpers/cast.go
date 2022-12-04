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
