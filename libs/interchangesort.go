package libs

import (
	"reflect"
)

// Big(O) = O(n*n)
func Intersort(arr []int) {
	swap := reflect.Swapper(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(i, j)
			}
		}
	}
}
