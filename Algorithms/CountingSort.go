package Algorithms

func CountingSort (arr []int, r int) []int {

	elems := make([]int, r)

	for i := range arr{
		elems[arr[i]]++
	}

	k := 0
	for i := range elems {
		if elems[i] != 0 {
			for j := 0; j < elems[i] ; j++ {
				arr[k] = i
				k++
			}
		}
	}
	return arr
}
