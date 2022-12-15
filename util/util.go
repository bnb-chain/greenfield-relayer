package util

func IndexOf(element string, data []string) int {
	for i, v := range data {
		if element == v {
			return i
		}
	}
	return -1
}
