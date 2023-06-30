package libs

import (
	"reflect"
)

// Big(O) = O(n*n)
func Bubblesort(arr []int) {
	swap := reflect.Swapper(arr)
	for i := 0; i < len(arr)-1; i++ {
		j := len(arr) - 1
		for j > i {
			if arr[j] < arr[j-1] {
				swap(j, j-1)
			}
			j--
		}
	}
}
