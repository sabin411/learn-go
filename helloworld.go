package main

import (
	"bufio"
	"fmt"
	"os"
)
var pl = fmt.Println


func main(){
	// learning variables
	var vName string = "Sabin shrestha";
	pl(vName, "is my name!!");
	pl("What is you name");
	reader := bufio.NewReader(os.Stdin);

	name, err :=reader.ReadString('\n');

	if(err != nil){
		pl("You got an error");
	}else{
		pl("hello", name)
	}
}