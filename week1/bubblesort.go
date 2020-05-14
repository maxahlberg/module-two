package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Promt user to input a array of upt to 10
	tenInts := make([]int, 10, 10)
	fmt.Printf("Input up to 10 comma separated integers, eg. '199,-3,5,2' (without the ' ) : \n")
	inputReader := bufio.NewScanner(os.Stdin)
	var txt string
	//adding them to the slice
	inputReader.Scan()
	txt = inputReader.Text()
	if len(txt) > 0 {
		values := strings.Split(txt, ",")
		var row []int
		for _, v := range values {
			fl, err := strconv.ParseFloat(strings.Trim(v, " "), 64)
			if err != nil {
				panic(fmt.Sprintf("Incorrect value for float64 '%v'", v))
			}
			row = append(row, int(fl))
		}
		tenInts = row[:10] //Takes the first ten inputs
		fmt.Printf("The input after parsing %v \n", tenInts)
	}
	// Create a bubble sort algorithm
	bubbleSort(&tenInts)
	// Print the sorted lice
	fmt.Printf("The Bubble Sorted slice is: %v \n", tenInts)

}

func bubbleSort(tenInts *[]int) {
	// Crreate a swap function to swap the values in the slice
	for i := 0; i < len(*tenInts); i++ {
		for ii := 0; ii < len(*tenInts)-1; ii++ {
			swap(*tenInts, ii)
		}
	}

}

func swap(tenInts []int, index int) {
	first := tenInts[index]
	next := tenInts[index+1]
	if first > next {
		tenInts[index] = next
		tenInts[index+1] = first
	}
}
