package Utils

import (
	"math/rand"
	"time"
)

// n - size, r - range
func GetUnsortedArray(n, r int) []int {
	arr := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < n; i++ {
		arr[i] = rand.Intn(r)
	}
	return arr
}
