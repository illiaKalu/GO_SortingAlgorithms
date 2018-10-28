package Algorithms

func SelectionSort (arr []int) []int {
	var max int;
	for i := len(arr) - 1; i > 1; i--  {
		max = i

		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[max] {
				max = j
			}
		}

		if i != max {
			arr[i], arr[max] = arr[max], arr[i]
		}

	}
	return arr
}
