package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap swaps two elements of a slice of integers
func Swap(numbers []int, index int) {
	numbers[index], numbers[index+1] = numbers[index+1], numbers[index]
}

// BubbleSort sorts a slice of integers
func BubbleSort(numbers []int) {
	n := len(numbers)
	var swapped bool
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers separated by space : ")
	input, _ := reader.ReadString('\n')
	stringNumbers := strings.Split(input, " ")
	numbers := []int{}
	for _, number := range stringNumbers {
		integerNumber, err := strconv.Atoi(strings.TrimSpace(number))
		if err == nil {
			numbers = append(numbers, integerNumber)
		}
	}
	BubbleSort(numbers)
	fmt.Print(numbers)
}
