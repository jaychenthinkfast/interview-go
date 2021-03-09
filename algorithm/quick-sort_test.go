package algorithm

import (
	"testing"
)

//	go test -run=TestQuickSort -v
func TestQuickSort(t *testing.T) {
	//=== RUN   TestQuickSort
	//quick-sort_test.go:11: [3 5 2 7 8 9]
	//quick-sort_test.go:13: [2 3 5 7 8 9]
	//--- PASS: TestQuickSort (0.00s)
	//PASS
	//ok      interview/algorithm     3.063s
	nums := []int{3, 5, 2, 7, 8, 9}
	t.Log(nums)
	QuickSort(nums, 0, len(nums)-1)
	t.Log(nums)
}
