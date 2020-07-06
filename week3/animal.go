package main

import (
	"fmt"
	"time"
)

type Animal struct {
	food  string
	loco  string
	sound string
}

func (a *Animal) Eat() {
	fmt.Printf(" eats: %v \n", a.food)
	fmt.Printf("------------------ Lets go again --------------\n")
	time.Sleep(2 * time.Second)
}

func (a *Animal) Move() {
	fmt.Printf(" goes from A to B by: %v \n", a.loco)
	fmt.Printf("------------------ Lets go again --------------\n")
	time.Sleep(2 * time.Second)
}

func (a *Animal) Speak() {
	fmt.Printf(" sound is: %v \n", a.sound)
	fmt.Printf("------------------ Lets go again --------------\n")
	time.Sleep(2 * time.Second)
}

func main() {
	// Defining animal
	m := make(map[string]Animal)
	cow := newAnimal("grass", "walk", "moo")
	bird := newAnimal("worms", "fly", "peep")
	snake := newAnimal("mice", "slither", "hsss")
	m["cow"] = cow
	m["bird"] = bird
	m["snake"] = snake
	var animal, information string
	for 1 == 1 {
		welcome()
		_, err := fmt.Scan(&animal, &information)
		if err != nil {
			fmt.Printf("failing reading input with scan\n")
		}
		if animal == "x" {
			break
		}
		animals := []string{"cow", "bird", "snake"}
		for _, v := range animals {
			if animal == v {
				any := m[animal]
				fmt.Printf("The %s", animal)
				switch {
				case information == "speak":
					any.Speak()
				case information == "move":
					any.Move()
				case information == "eat":
					any.Eat()
				default:
					fmt.Printf("no match in switch\n")
				}
			}
		}
	}
}

func newAnimal(eat, move, speak string) (a Animal) {
	a.food = eat
	a.loco = move
	a.sound = speak
	return a
}

func welcome() {
	fmt.Printf("Welcome to Animal info: \n" +
		"Ask me anout any animal, as long as it is either 'cow', 'bird' or 'snake' \n" +
		"Ask me then what it can 'eat', how it 'move' or 'speak'\n" +
		"for example 'bird eat'\n" +
		"write 'x x' to exit\n" +
		">\n")
}
