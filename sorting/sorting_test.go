package sorting

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{13, 46, 24, 52, 20, 9}
	sortedArray := []int{9, 13, 20, 24, 46, 52}
	n := len(arr) // length of array
	// expectedSelectionSortResult := []int{9, 13, 20, 24, 46, 52}
	// gotSelectionSortResult := SelectionSort(arr, n)
	// SelectionSort(arr, n)
	// got := bubbleSort(arr, n)
	got := insertionSort(arr, n)
	fmt.Println(got)
	compareArray(sortedArray, got)
	assert.Assert(t, compareArray(got, sortedArray), true)
}

func compareArray(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
