package libs

import "reflect"

// Big(O) = O(n*n)
func Selectsort(arr []int) {
	swap := reflect.Swapper(arr)
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				min = j
			}
		}
		if i != min {
			swap(i, min)
		}

	}
}
