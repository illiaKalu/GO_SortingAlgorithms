package Algorithms

func OriginalMergeSort (arr []int) []int {

	var tmp = make([]int, len(arr))

	var l, r = 0, len(arr) - 1
	sortO(arr, tmp, l, r)

	return arr
}
func sortO(arr, tmp []int, l, r int) {

	if l < r {
		center := (l + r) / 2
		// sort left
		sortO(arr, tmp, l, center)
		// sort right
		sortO(arr, tmp, center + 1, r)
		mergeO(arr, tmp, l, center + 1, r)
	}
}

func mergeO(arr, tmp []int, left, right, rightEnd int) {
	leftEnd := right - 1
	k := left
	num := rightEnd - left + 1

	for left <= leftEnd && right <= rightEnd {
		if arr[left] <= arr[right] {
			tmp[k] = arr[left]
			left++
		} else {
			tmp[k] = arr[right]
			right++
		}
		k++
	}

	for left <= leftEnd {
		tmp[k] = arr[left]
		k++
		left++
	}

	for right <= rightEnd {
		tmp[k] = arr[right]
		k++
		right++
	}

	for i := 0; i < num; i++ {
		arr[rightEnd] = tmp[rightEnd]
		rightEnd--
	}

}
