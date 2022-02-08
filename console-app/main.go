package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Pure Black"
	coffees[2] = "Americano"
	coffees[3] = "Iced Latte"
	coffees[4] = "Iced Black"
	coffees[5] = "Mad Mocha"
	coffees[0] = "to quit the program"

	fmt.Println("Menu")
	fmt.Println("----")
	fmt.Println("1 - Pure Black")
	fmt.Println("2 - Americano")
	fmt.Println("3 - Iced Latte")
	fmt.Println("4 - Iced Black")
	fmt.Println("5 - Mad Mocha")
	fmt.Println("Q - Quit the program")
	fmt.Printf("\n")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		// converts a rune that has been cast into a string into an integer
		i, _ := strconv.Atoi(string(char))

		// uses the integer value we got from the rune to get the map element
		// TODO: check if the character is within the desired range before printing
		fmt.Printf("You chose %s.\n", coffees[i])

		if char == 'q' || char == 'Q' {
			break
		}

	}
}
