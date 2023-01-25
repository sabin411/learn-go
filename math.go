package main

import (
	"fmt"
	"math/rand"
	"time"
)
var pl = fmt.Println

func main(){
	seedSecs := time.Now().Unix();
	rand.Seed(seedSecs);
	randNum := rand.Intn(50) + 1;
	pl("My favorite number is", randNum);
	// Things to keep in mind 
	// 1. There are a lot of math function that i will rnd when required. 
}
