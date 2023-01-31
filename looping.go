package main

import (
	"fmt"
	"time"
	"math/rand"
)
var pl = fmt.Println

func main(){
	seedSec := time.Now().Unix();
	rand.Seed(seedSec);
	randNum := rand.Intn(50) + 1

	for true {
		pl("Guess a number between 1 to 50");
		var input int;
		fmt.Scanln(&input);
		if input == randNum {
			pl("You guessed it right!");
			break;
		}
		if input < randNum {
			pl("Too low!");
		}
		if input > randNum {
			pl("Too high!");
		}
	}
}