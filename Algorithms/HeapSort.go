package Algorithms

func HeapSort (arr []int) []int {

	for i := len(arr) / 2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}
	for i:= len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
	return arr
}

func heapify(arr []int, n, i int)  {

	biggest := i
	l := i * 2 + 1
	r := i * 2 + 2

	if l < n && arr[l] > arr[biggest] {
		biggest = l
	}

	if r < n && arr[r] > arr[biggest] {
		biggest = r
	}

	if biggest != i {
		arr[i], arr[biggest] = arr[biggest], arr[i]
		heapify(arr, n, biggest)
	}
}
