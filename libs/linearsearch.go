package libs

import (
	"fmt"
)

// Big(O) = O(n)
func Lsearch(arr []int, x int) {
	flag := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			flag = true
			fmt.Println("Found at position ", i)
			break
		}
	}
	if flag == false {
		fmt.Println("Not found")
	}
}
