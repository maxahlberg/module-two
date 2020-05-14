package main

import "fmt"

func main() {
	// Promt user to input a array of upt to 10
	tenInts := []int{4, 5, 1, 2, -2}

	// Create a bubble sort algorithm
	bubbleSort(&tenInts)
	// Print the sorted lice
	fmt.Printf("The Bubble Sorted slice is: %v \n", tenInts)

}

func bubbleSort(tenInts *[]int) {
	// Crreate a swap function to swap the values in the slive
	for i := 0; i < len(*tenInts); i++{
		for ii := 0; ii < len(*tenInts)-1; ii++ {
			swap(*tenInts, ii, i)
		}
	}

}

func swap(tenInts []int, index, round int){
	first := tenInts[index]
	next := tenInts[index+1]
	if first > next {
		tenInts[index] = next
		tenInts[index+1] = first
	}
	fmt.Printf("The ten ints to be swapped %v, index: %v, round in for loop: %v \n", tenInts, index, round)
}
