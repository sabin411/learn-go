package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println;

func main(){
	string1 := "A new world!";
	replacer := strings.NewReplacer("A", "Another");
	string2 := replacer.Replace(string1);
	pl(string2, "It has been replaced!!");
	pl("The string contains Another", strings.Contains(string2, "Another"));
	pl("Length of string2", len(string2));
}