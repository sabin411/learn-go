package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main(){
	now := time.Now();
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}