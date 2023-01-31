package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)
var pl = fmt.Println

func main(){
	seedSecs := time.Now().Unix();
	rand.Seed(seedSecs);
	randNum := rand.Intn(50) + 1;
	pl("My favorite number is", randNum);
	pl("Max value", math.Max(34, 33))
	pl("Max value", math.Min(323, 2))
	// Things to keep in mind 
	// 1. There are a lot of math function that i will rnd when required. 

	// formatted datatypes 
	// %d - int
	// %f - float
	// %t - bool
	// %s - string
	// %c - char
	// %v - value
	// %T - type
	// %b - binary
	// %x - hex
	// %e - scientific notation
	// %p - pointer
	fmt.Printf("%9v %T %b %x %e %p", 3.14, true, "sabin", "#000");


}
