package sorting

import "fmt"

func SelectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		var min int = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
			// swap(&arr[i], &arr[j])
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	fmt.Println("inside selection sort", arr)
}
