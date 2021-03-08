package algorithm

import (
	"testing"
)

//	go test -run=TestSelectionSort -v
func TestSelectionSort(t *testing.T) {
	//=== RUN   TestSelectionSort
	//selection-sort_test.go:10: [3 5 2 7 8 9]
	//selection-sort_test.go:12: [2 3 5 7 8 9]
	//--- PASS: TestSelectionSort (0.00s)
	//PASS
	//ok      interview/algorithm     4.002s
	nums := []int{3, 5, 2, 7, 8, 9}
	t.Log(nums)
	numsSorted := SelectionSort(nums)
	t.Log(numsSorted)
}
