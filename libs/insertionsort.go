package libs

func Insertsort(arr []int) {
	position, x := 0, 0
	for i := 0; i < len(arr)-1; i++ {
		x = arr[i+1]
		position = i
		for position >= 0 && arr[position] > x {
			arr[position+1] = arr[position]
			position--
		}

		arr[position+1] = x
	}
}
