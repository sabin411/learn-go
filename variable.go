package main

import (
	"fmt"
	"strings"
)
var pl = fmt.Println;

func main(){
	string1 := "A world"
	replacer := strings.NewReplacer("A", "Another");
	string2 := replacer.Replace(string1);

	pl(string2, "Print this");

	pl("Length", len(string1));
	pl("Container another: ", strings.Contains(string2, "Another"));
}