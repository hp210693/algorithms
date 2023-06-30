package libs

import "fmt"

// Big(O) = log(n)
func Bsearch(arr []int, x int) {
	flag := false
	left, mid := 0, 0
	right := len(arr) - 1
	for left <= right {
		mid = (left + right) / 2
		if arr[mid] == x {
			flag = true
			fmt.Println("Found at position ", mid)
			break
		}
		if arr[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if flag == false {
		fmt.Println("Not found")
	}
}

// Recursion
// Big(O) = log(n)
func BRsearch(arr []int, x, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return BRsearch(arr, x, left, mid-1)
		} else {
			return BRsearch(arr, x, mid+1, right)
		}
	}
	return -1
}
