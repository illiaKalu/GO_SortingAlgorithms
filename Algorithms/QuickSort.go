package Algorithms

func QuickSort (arr []int) []int {

	quickSort(arr, 0, len(arr) - 1)
	return arr
}

func quickSort(arr []int, l, r int)  {
	if l < r {
		p := part(arr, l, r)

		quickSort(arr, l, p - 1)
		quickSort(arr, p + 1, r)
	}
}

func part(arr []int, l, r int)int {
	pivot := arr[r]
	i := l - 1

	for j := l; j < r; j++ {
		if arr[j] <= pivot {
			i++
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	arr[i + 1], arr[r] = arr[r], arr[i + 1]

	return i + 1
}
