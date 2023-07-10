package sorting

// Big(O) = O(n*n)
func SelectionSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// swap elements
		arr[i], arr[min] = arr[min], arr[i]
	}
}
