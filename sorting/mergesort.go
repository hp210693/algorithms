package sorting

// Big(O) = O(nlog(n))
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	return merge(MergeSort(arr[0:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

func merge(left, right []int) []int {
	temp := make([]int, len(left)+len(right))

	for i, j, k := 0, 0, 0; k < len(temp); k++ {
		if i > len(left)-1 && j < len(right) {
			temp[k] = right[j]
			j++
			continue
		}
		if j > len(right)-1 && i < len(left) {
			temp[k] = left[i]
			i++
			continue
		}
		if left[i] < right[j] {
			temp[k] = left[i]
			i++
		} else {
			temp[k] = right[j]
			j++
		}
	}
	return temp
}
