package main

import (
	"fmt"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (a animal) Eat() string {
	return a.food
}

func (a animal) Speak() string {
	return a.noise
}

func (a animal) Move() string {
	return a.locomotion
}

func main() {
	var cow, snake, bird animal
	var usrAnimal, usrAction string
	fmt.Println("Welcome to this program!")
	fmt.Println("To use this program, type an animals name and action")
	for {
		fmt.Print(">")
		fmt.Scan(&usrAnimal, &usrAction)
		fmt.Println(usrAnimal, usrAction)
		usrAnimal = strings.ToLower(usrAnimal)
		usrAction = strings.ToLower(usrAction)
		switch {
		case usrAnimal == "snake":
			snake.food = "mice"
			snake.locomotion = "slither"
			snake.noise = "hssss"
			if usrAction == "move" {
				fmt.Println(snake.Move())
			} else if usrAction == "eat" {
				fmt.Println(snake.Eat())
			} else if usrAction == "speak" {
				fmt.Println(snake.Speak())
			} else {
				fmt.Println("Action not recognized!")
			}
		case usrAnimal == "bird":
			bird.food = "worms"
			bird.locomotion = "fly"
			bird.noise = "peep"
			if usrAction == "move" {
				fmt.Println(bird.Move())
			} else if usrAction == "eat" {
				fmt.Println(bird.Eat())
			} else if usrAction == "speak" {
				fmt.Println(bird.Speak())
			} else {
				fmt.Println("Action not recognized!")
			}
		case usrAnimal == "cow":
			cow.food = "grass"
			cow.locomotion = "walk"
			cow.noise = "moo"
			if usrAction == "move" {
				fmt.Println(cow.Move())
			} else if usrAction == "eat" {
				fmt.Println(cow.Eat())
			} else if usrAction == "speak" {
				fmt.Println(cow.Speak())
			} else {
				fmt.Println("Action not recognized!")
			}
		default:
			fmt.Println("animal not recognized!")
		}
	}
}
