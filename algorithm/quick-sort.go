package algorithm

func Partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := Partition(a, lo, hi)
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}
