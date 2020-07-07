package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	food  string
	loco  string
	sound string
}

func (c cow) Eat()   { fmt.Printf("%s\n", c.food) }
func (c cow) Move()  { fmt.Printf("%s\n", c.loco) }
func (c cow) Speak() { fmt.Printf("%s\n", c.sound) }

type bird struct {
	food  string
	loco  string
	sound string
}

func (b bird) Eat()   { fmt.Printf("%s\n", b.food) }
func (b bird) Move()  { fmt.Printf("%s\n", b.loco) }
func (b bird) Speak() { fmt.Printf("%s\n", b.sound) }

type snake struct {
	food  string
	loco  string
	sound string
}

func (s snake) Eat()   { fmt.Printf("%s\n", s.food) }
func (s snake) Move()  { fmt.Printf("%s\n", s.loco) }
func (s snake) Speak() { fmt.Printf("%s\n", s.sound) }

func executeNewAnimal(name, animalType string) {
	switch {
	case animalType == "cow":
		AnimalMap[name] = cow{
			food:  "grass",
			loco:  "walk",
			sound: "moo",
		}

	case animalType == "bird":
		AnimalMap[name] = bird{
			food:  "worms",
			loco:  "fly",
			sound: "peep",
		}
	case animalType == "snake":
		AnimalMap[name] = snake{
			food:  "mice",
			loco:  "slither",
			sound: "hsss",
		}
	default:
		fmt.Printf("The animal type did not match 'cow', 'bird' or 'snake', which is has to\n")
	}
}

func answerQuery(name, infoRequested string) {
	for key, value := range AnimalMap {
		if name == key {
			switch {
			case infoRequested == "eat":
				fmt.Printf("The animal %s eats ", name)
				value.Eat()
			case infoRequested == "move":
				fmt.Printf("The animal %s ", name)
				value.Move()
			case infoRequested == "speak":
				fmt.Printf("The animal %s says ", name)
				value.Speak()
			default:
				fmt.Printf("The info requested did not match 'eat', 'move' or 'speak'\n")
			}
		}
	}
}

var AnimalMap map[string]Animal

func main() {
	AnimalMap = make(map[string]Animal)
	var input1, input2, input3 string
	for {
		fmt.Printf(">")
		_, err := fmt.Scan(&input1, &input2, &input3)
		if err != nil {
			fmt.Printf("Error scaning input %s \n", err)
		}
		switch {
		case input1 == "newanimal":
			fmt.Printf("Created it!\n")
			name := input2
			animalType := input3
			executeNewAnimal(name, animalType)
		case input1 == "query":
			name := input2
			infoRequested := input3
			answerQuery(name, infoRequested)
		default:
			fmt.Printf("The first input argument did not match either 'newanimal' or 'query', " +
				"which it has to. Try again\n")
		}
	}
}
