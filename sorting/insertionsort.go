package sorting

// Big(O) = O(n*n)
func InserttionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		position := i
		for position >= 1 && arr[position-1] > x {
			arr[position] = arr[position-1]
			position--
		}

		arr[position] = x
	}
}
