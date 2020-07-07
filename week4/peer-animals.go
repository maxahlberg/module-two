package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type AnimalType struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func NewAnimal(x string, model string) AnimalType {
	if model == "cow" {
		return AnimalType{name: x, food: "grass", locomotion: "walk", noise: "moo"}
	}
	if model == "bird" {
		return AnimalType{name: x, food: "worms", locomotion: "fly", noise: "peep"}
	}
	if model == "snake" {
		return AnimalType{name: x, food: "mice", locomotion: "slither", noise: "hsss"}
	} else {
		return AnimalType{name: x, food: "NA", locomotion: "NA", noise: "NA"}
	}
}

func (ani AnimalType) Eat() {
	fmt.Println(ani.food)
}

func (ani AnimalType) Move() {
	fmt.Println(ani.locomotion)
}

func (ani AnimalType) Speak() {
	fmt.Println(ani.noise)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sli := make([]string, 3)
	var animalSli []AnimalType

	cow := AnimalType{name: "cow", food: "grass", locomotion: "walk", noise: "moo"}
	bird := AnimalType{name: "bird", food: "worms", locomotion: "fly", noise: "peep"}
	snake := AnimalType{name: "snake", food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		fmt.Println("Enter a command (newanimal or query).")
		fmt.Printf(">")
		scanner.Scan()
		sli = strings.Fields(scanner.Text())

		if sli[0] == "query" {
			if sli[1] == "cow" {
				if sli[2] == "eat" {
					cow.Eat()
				}
				if sli[2] == "move" {
					cow.Move()
				}
				if sli[2] == "speak" {
					cow.Speak()
				}
			}

			if sli[1] == "bird" {
				if sli[2] == "eat" {
					bird.Eat()
				}
				if sli[2] == "move" {
					bird.Move()
				}
				if sli[2] == "speak" {
					bird.Speak()
				}
			}

			if sli[1] == "snake" {
				if sli[2] == "eat" {
					snake.Eat()
				}
				if sli[2] == "move" {
					snake.Move()
				}
				if sli[2] == "speak" {
					snake.Speak()
				}
			} else {
				for j := range animalSli {
					if animalSli[j].name == sli[1] {
						if sli[2] == "eat" {
							fmt.Println(animalSli[j].food)
						}
						if sli[2] == "move" {
							fmt.Println(animalSli[j].locomotion)
						}
						if sli[2] == "speak" {
							fmt.Println(animalSli[j].noise)
						}
					}
				}
			}
		} else if sli[0] == "newanimal" {
			for i := 0; i < 100; i++ {
				newThing := NewAnimal(sli[1], sli[2])
				animalSli = append(animalSli, newThing)
				break
			}
			fmt.Println("Created it!")
		} else {
			fmt.Println("Invalid input. Try again.")
			fmt.Println(sli[0])
		}
	}
}
