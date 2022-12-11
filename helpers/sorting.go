package helpers

func Reverse(arrayToReverse []int) []int {
	reversedArray := []int{}

	for i := range arrayToReverse {
		reversedArray = append(reversedArray, arrayToReverse[len(arrayToReverse)-1-i])
	}
	return reversedArray
}
