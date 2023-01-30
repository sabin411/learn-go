package main

import (
	"fmt"
)
var pl = fmt.Println

func main(){
	myArr := [5]int{1, 2, 3, 4, 5}
	for i, v := range myArr {
		pl(i, v)
	}
	a := [4]int{1, 2, 4, 5}
	a[3] = 400
	pl(a)

	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
