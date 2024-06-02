package sorting

func insertionSort(arr []int, n int) []int {
	for i := 0; i <= n-1; i++ {
		j := i
		for ; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
	return arr
}
