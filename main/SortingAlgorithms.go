package main

import (
	"SortingAlgorithms/Algorithms"
	"time"
	"log"
	"SortingAlgorithms/Utils"
	"fmt"
)


func main() {
	var n, r = 20000, 10000

	start := time.Now()
	unsortedArray := Utils.GetUnsortedArray(n, r)
	unsortedArray2 := Utils.GetUnsortedArray(n, r)
	unsortedArray3 := Utils.GetUnsortedArray(n, r)
	unsortedArray4 := Utils.GetUnsortedArray(n, r)
	unsortedArray5 := Utils.GetUnsortedArray(n, r)
	unsortedArray6 := Utils.GetUnsortedArray(n, r)
	unsortedArray7 := Utils.GetUnsortedArray(n, r)
	elapsed := time.Since(start)
	//log.Printf("Creation took %s", elapsed)

	start = time.Now()
	Algorithms.BubbleSort(unsortedArray)
	elapsed = time.Since(start)
	log.Printf("bubble took %s", elapsed)

	start = time.Now()
	Algorithms.InsertionSort(unsortedArray2)
	elapsed = time.Since(start)
	log.Printf("insertion took %s", elapsed)

	start = time.Now()
	Algorithms.OriginalMergeSort(unsortedArray3)
	elapsed = time.Since(start)
	log.Printf("merge took %s", elapsed)

	start = time.Now()
	Algorithms.CountingSort(unsortedArray4, r)
	elapsed = time.Since(start)
	log.Printf("count took %s", elapsed)

	start = time.Now()
	Algorithms.SelectionSort(unsortedArray5)
	elapsed = time.Since(start)
	log.Printf("selection took %s", elapsed)

	start = time.Now()
	Algorithms.HeapSort(unsortedArray6)
	elapsed = time.Since(start)
	log.Printf("heap took %s", elapsed)

	start = time.Now()
	Algorithms.QuickSort(unsortedArray7)
	elapsed = time.Since(start)
	log.Printf("quick took %s", elapsed)

}

func printArr(ints []int) {
	for _, v := range ints {
		fmt.Print(v, ", ")
	}
}
