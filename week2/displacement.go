package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// promt the user to enter a_0, v_0 and s_0
	fmt.Printf("enter initial acceleration, velocity, displacement. Comma separated ex. 3,2,4 : \n")
	inputReader := bufio.NewScanner(os.Stdin)
	//adding them to the slice
	inputedStates := parseInput(inputReader)
	a_0 := inputedStates[0]
	v_0 := inputedStates[1]
	s_0 := inputedStates[2]
	// promt uder to enter a time, t.
	fmt.Printf("Enter a time: \n")
	inputedStates = parseInput(inputReader)
	t := inputedStates[0]
	// call the general displacement func
	fn := genDisplaceFn(a_0, v_0, s_0)
	fmt.Printf("The Displacement is %v \n", fn(t))
}

func genDisplaceFn(a, v, s float64) func(time float64)(disp float64)  {
	return func(t float64) (disp float64){
		fmt.Printf("all initail variables: a: %v, v: %v, s:%v, t:%v, \n", a, v, s, t)
		disp = 0.5*a*math.Pow(t, 2)+ v*t + s
		return disp
	}
}

func parseInput(inputReader *bufio.Scanner) []float64{
	initialCond := make([]float64, 0, 3)
	var txt string
	inputReader.Scan()
	txt = inputReader.Text()
	if len(txt) > 0 {
		values := strings.Split(txt, ",")
		var row []float64
		for _, v := range values {
			fl, err := strconv.ParseFloat(strings.Trim(v, " "), 64)
			if err != nil {
				panic(fmt.Sprintf("Incorrect value for float64 '%v'", v))
			}
			row = append(row, fl)
		}
		initialCond = row
	}
	return initialCond
}
