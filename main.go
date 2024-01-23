package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().Unix())
	sourced := rand.New(source)
	var secret_number int = sourced.Intn(10)
	fmt.Println("Please guess a number")
	for {
		var i int
		_, _ = fmt.Scanf("%d", &i)
		if i == secret_number {
			fmt.Println("You Guess right!")
			break
		} else if i < secret_number {
			fmt.Println("You Guess below the secret number!")
		} else {
			fmt.Println("You Guess above the secret number!")
		}

	}

}
