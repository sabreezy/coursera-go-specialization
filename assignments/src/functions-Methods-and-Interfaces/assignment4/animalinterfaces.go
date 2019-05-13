package main

import (
	"fmt"
	"strings"
)

type animal interface {
	speak()
	move()
	eat()
}

type bird struct {
	food       string
	locomotion string
	sound      string
}

type cow struct {
	food       string
	locomotion string
	sound      string
}

type snake struct {
	food       string
	locomotion string
	sound      string
}

func (b bird) speak() {
	fmt.Println("peep")
}

func (b bird) eat() {
	fmt.Println("worms")
}

func (b bird) move() {
	fmt.Println("fly")
}

func (s snake) speak() {
	fmt.Println("hssss")
}

func (s snake) move() {
	fmt.Println("slither")
}

func (s snake) eat() {
	fmt.Println("mice")
}

func (c cow) speak() {
	fmt.Println("moo")
}

func (c cow) eat() {
	fmt.Println("grass")
}

func (c cow) move() {
	fmt.Println("walk")
}

func main() {
	m := make(map[string]animal)
	fmt.Println("Welcome to this program")
	for {
		fmt.Printf(">\n")
		var command, animalName, animalType, commandType string
		fmt.Scanln(&command, &animalName, &animalType)
		command = strings.ToLower(command)
		animalName = strings.ToLower(animalName)
		animalType = strings.ToLower(animalType)
		//first check if the command is animal or query
		for !(command == "newanimal") && !(command == "query") {
			fmt.Println("Not valid keyword! Please use either newanimal or query!")
			fmt.Scan(&command)
		}
		if command == "newanimal" {
			for !(animalType == "cow") && !(animalType == "snake") && !(animalType == "bird") {
				fmt.Println("Incorrect animal type! Please enter only bird, snake, or cow!")
				fmt.Scan(&animalType)
				animalType = strings.ToLower(animalType)
			}
		}
		if command == "query" {
			commandType = animalType
			for !(commandType == "move") && !(commandType == "eat") && !(commandType == "speak") {
				fmt.Println("incorrect animal command! Please input only move,eat, or speak!")
			}
		}
		// to store animal names and types
		switch {
		case command == "newanimal":
			switch {
			case animalType == "snake":
				a1 := snake{}
				m[animalName] = a1
				fmt.Println("Created it!")
			case animalType == "cow":
				a2 := cow{}
				m[animalName] = a2
				fmt.Println("Created it!")
			case animalType == "bird":
				a3 := bird{}
				m[animalName] = a3
				fmt.Println("Created it!")
			}
		case command == "query":
			if val, ok := m[animalName]; ok {
				fmt.Println("found animal name in map!")
				switch {
				case commandType == "move":
					val.move()
					fmt.Println("performed action!")
				case commandType == "eat":
					val.eat()
					fmt.Println("performed action!")
				case commandType == "speak":
					val.speak()
					fmt.Println("performed action!")
				}
			}
		}
	}
}
