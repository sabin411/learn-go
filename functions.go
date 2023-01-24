package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)
var pl = fmt.Println

// Returns sum of values
	func getSum(x int, y int) int {
		return (x + y)
	}

	func sayHello() {
		pl("Hello")
	}

	// Return multiple values
	func getTwo(x int) (int, string) {
		return x + 1, strconv.Itoa(x);
	}

func main(){
	var total = getSum(30, 40);
	pl("The total sum is: ",total)
	sayHello();
	pl(getTwo(2));

	// runes datatype runes are unicodes that represents characters 
	rStr := "sabin stha";
	pl("Rune code count: ", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal);
	}

}