package main

import (
	"fmt"
)

var pl = fmt.Println;

func changeString (str *string){
	*str = "Changed";
}

func changeString2 (str string)(changedStr string){
	str = "Changed";
	return str;
}

func main() {
	// pointer and dereferance
	x := 10;
	y := &x; // y is memory location of value 10. 
	pl(x, y);
	*y = 30; // here dereferencing pointer
	pl(x, y);
	x = 50;
	pl(x, y);
	
	myStr := "Hello world"
	// Using pointer.
	pl(myStr);
	changeString(&myStr); // Getting the pointer of myStr
	pl(myStr);

	// Without using pointer.
	myStr2 := "Hello world"
	pl(myStr2);
	myStr2 = changeString2(myStr2); // Getting the changed value from the function
	pl(myStr2);
}