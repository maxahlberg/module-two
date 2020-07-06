package main

import (
	"fmt"
	"time"
)

type Animal struct {
	sort  string
	food  string
	loco  string
	sound string
}

func (a *Animal) Eat() {
	fmt.Printf("----------------- Your answer is --------------\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("The %s eats: %v \n", a.sort, a.food)
	fmt.Printf("------------------ Lets go again --------------\n")
	time.Sleep(2 * time.Second)
}

func (a *Animal) Move() {
	fmt.Printf("----------------- Your answer is --------------\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("The %s goes from A to B by: %v \n", a.sort, a.loco)
	fmt.Printf("------------------ Lets go again --------------\n")
	time.Sleep(2 * time.Second)
}

func (a *Animal) Speak() {
	fmt.Printf("----------------- Your answer is --------------\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("The %s sound is: %v \n", a.sort, a.sound)
	fmt.Printf("------------------ Lets go again --------------\n")
	time.Sleep(2 * time.Second)
}

func main() {
	// Defining animal
	Cow := newAnimal("cow", "grass", "walk", "moo")
	Bird := newAnimal("bird", "worms", "fly", "peep")
	Snake := newAnimal("snake", "mice", "slither", "hsss")
	var animal, information string
	for 1 == 1 {
		welcome()
		_, err := fmt.Scan(&animal, &information)
		if err != nil {
			fmt.Printf("failing reading input with scan\n")
		}
		if animal == "x"{break}
		animals := []Animal{Cow, Bird, Snake}
		for _, v := range animals {
			if animal == v.sort {
				switch {
				case information == "speak":
					v.Speak()
				case information == "move":
					v.Move()
				case information == "eat":
					v.Eat()
				default:
					fmt.Printf("no match in switch\n")
				}
			}
		}
	}
}

func newAnimal(sort, eat, move, speak string) (a Animal) {
	a.sort = sort
	a.food = eat
	a.loco = move
	a.sound = speak
	return a
}
func welcome(){
	fmt.Printf("Welcome to Animal info: \n" +
		"Ask me anout any animal, as long as it is either 'cow', 'bird' or 'snake' \n" +
		"Ask me then what it can 'eat', how it 'move' or 'speak'\n" +
		"for example 'bird eat'\n" +
		"write 'x x' to exit\n"+
		">\n")
}
