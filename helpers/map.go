package helpers

func AppendAll(map1, map2 map[string]int) map[string]int {
	var totalMap = map[string]int{}
	for index, element := range map1 {
		totalMap[index] = element
	}
	for index, element := range map2 {
		totalMap[index] = element
	}
	return totalMap
}
