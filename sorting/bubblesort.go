package sorting

// Big(O) = O(n*n)
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Big(O) = O(n)
func BubbleSortImproved(arr []int) {
	swapped := true
	for i := len(arr) - 1; i >= 0 && swapped; i-- {
		swapped = false
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
	}
}
