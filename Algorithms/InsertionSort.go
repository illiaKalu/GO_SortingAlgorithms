package Algorithms

func InsertionSort (arr []int) []int {
	var j int
	for i:= 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			j = i - 1
			for j >= 0 && arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				j--
			}
		}
	}
	return arr
}


