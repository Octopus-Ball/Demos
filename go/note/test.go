package main


func Sum(set []int) int {
	var result int
	for _, num := range set {
		result += num
	}
	return result
}