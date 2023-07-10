package main

import (
	"algorithms/examples"
	"algorithms/searching"
	"algorithms/sorting"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func autogenarr() []int {
	arr := []int{}
	source := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(source)
	for i := 0; i < 50; i++ {
		arr = append(arr, randSource.Intn(100))
	}
	return arr
}

func m_lissajous(o *os.File) {
	examples.Lissajous(o)
}
func m_fetchurl() {
	examples.FetchURL()
}
func m_fetchallurl() {
	examples.FetchAllURL()
}
func m_webserver() {
	examples.Webserver()
}

func main() {
	//var choosen string

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+\t|+++Choosen funtions:\t\t\t\t+")
	fmt.Println("+\t |I.SORTTING\t\t\t\t\t+")
	fmt.Println("+\t |---1.1.Algorithms bubble sort\t\t\t+")
	fmt.Println("+\t |---1.2.Algorithms bubble sort improved\t+")
	fmt.Println("+\t |---1.3.Algorithms selection sort\t\t+")
	fmt.Println("+\t |---1.4.Algorithms insertion sort\t\t+")
	fmt.Println("+\t |---1.5.Algorithms shell sort\t\t\t+")
	fmt.Println("+\t |---1.6.Algorithms interchange sort\t\t+")
	fmt.Println("+\t |---1.7.Algorithms merge sort\t\t\t+")
	fmt.Println("+\t |II.Searching\t\t\t\t\t+")
	fmt.Println("+\t |---2.1.Algorithms binary seach\t\t+")
	fmt.Println("+\t |---2.2.Algorithms recursion binary seach\t+")
	fmt.Println("+\t |---2.3.Algorithms linear search\t\t+")
	fmt.Println("+\t |III.Examples\t\t\t\t\t+")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Print("Input your choosen = ")
	//fmt.Scan(&choosen)
	arr := autogenarr()
	//arr := []int{6, 5, 11, 10, 9, 1}
	x := 4
	fmt.Println("Before arr = ", arr)
	switch "1.7" {
	case "1.1":
		sorting.BubbleSort(arr)
	case "1.2":
		sorting.BubbleSortImproved(arr)
	case "1.3":
		sorting.SelectionSort(arr)
	case "1.4":
		sorting.InserttionSort(arr)
	case "1.5":
		fmt.Println("Comming soon")
	case "1.6":
		sorting.InterchangeSort(arr)
	case "1.7":
		fmt.Printf("Aeter arr  = %v\n", sorting.MergeSort(arr))
	case "2.1":
		searching.BSearch(arr, x)
	case "2.2":
		position := searching.BRSearch(arr, x, 0, len(arr))
		if position == -1 {
			fmt.Println("Not found")
		} else {
			fmt.Println("Found at position ", position)
		}
	case "2.3":
		searching.LinearSearch(arr, x)
	case "8":
		m_fetchurl()
	case "9":
		m_fetchallurl()
	case "10":
		m_webserver()
	}
	//fmt.Println("After arr  = ", arr)
}
