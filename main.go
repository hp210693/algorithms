package main

import (
	"algorithms/libs"
	"fmt"
	"math/rand"
	"time"
)

func m_bsearch(arr []int, x int) {
	libs.BRsearch(arr, x, 0, len(arr)-1)

	position := libs.BRsearch(arr, x, 0, len(arr)-1)
	if position == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found at position ", position)
	}
}

func m_linearsearch(arr []int, x int) {
	libs.Lsearch(arr, x)
}

func m_bubblesort(arr []int) {
	libs.Bubblesort(arr)
}

func m_insertionsort(arr []int) {
	libs.Insertsort(arr)
}

func m_intersort(arr []int) {
	libs.Intersort(arr)
}
func m_selectsort(arr []int) {
	libs.Selectsort(arr)
}

func autogenarr() []int {
	arr := []int{}
	source := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(source)
	for i := 0; i < 20; i++ {
		arr = append(arr, randSource.Intn(100))
	}
	return arr
}

func main() {
	var choosen int

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+\t|+++Choosen funtions:\t\t\t\t+")
	fmt.Println("+\t|---1.Algorithms binary seach\t\t\t+")
	fmt.Println("+\t|---2.Algorithms bubble sort\t\t\t+")
	fmt.Println("+\t|---3.Algorithms insertion sort\t\t\t+")
	fmt.Println("+\t|---4.Algorithms interchange sort\t\t+")
	fmt.Println("+\t|---5.Algorithms linear search\t\t\t+")
	fmt.Println("+\t|---6.Algorithms selection sort\t\t\t+")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Print("Input your number = ")
	fmt.Scan(&choosen)
	arr := autogenarr()
	x := 4
	fmt.Println("Before arr = ", arr)
	switch choosen {
	case 1:
		m_bsearch(arr, x)
	case 2:
		m_bubblesort(arr)
	case 3:
		m_insertionsort(arr)
	case 4:
		m_intersort(arr)
	case 5:
		m_linearsearch(arr, x)
	case 6:
		m_selectsort(arr)
	}
	fmt.Println("After arr  = ", arr)
}
