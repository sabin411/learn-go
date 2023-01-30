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

		func sumArray(arr []int) (int){
		sum := 0
		for _, value := range arr{
			sum += value
		}
		return sum
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

	// Arrays
	myArr := []int{1,2,3,4,5,6,7,8,9,10};
	pl("Sum of array: ", sumArray(myArr));

}