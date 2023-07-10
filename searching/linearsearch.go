package searching

import (
	"fmt"
)

// Big(O) = O(n)
func LinearSearch(arr []int, x int) {
	flag := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			flag = true
			fmt.Println("Found at position ", i)
			break
		}
	}
	if !flag {
		fmt.Println("Not found")
	}
}
