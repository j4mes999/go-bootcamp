package main

import (
	"fmt"
)

func main() {

	fmt.Println("****** Challenge 1 ******")
	superHeroes := []string{"spiderman", "batman", "cyclop", "joker", "Storm", "wonder woman"}
	challenge1(superHeroes)

}

func challenge1(superHeroes []string) {
	var marvelHeroes []string
	var dcHeroes []string

	for i, heroe := range superHeroes {
		switch {
		case i == 0:
			marvelHeroes = append(marvelHeroes, heroe)
		case i%2 == 0: //If i is even then is a marvel heroe
			marvelHeroes = append(marvelHeroes, heroe)
		case i%2 == 1: // if i is odd is DC
			dcHeroes = append(dcHeroes, heroe)
		}
	}
	fmt.Println("Marvel Heroes --> ", marvelHeroes)
	fmt.Println("DC Heroes --> ", dcHeroes)
}
