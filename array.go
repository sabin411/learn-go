package main

import (
	"fmt"
)
var pl = fmt.Println
var pf = fmt.Printf

func main(){
	myArr := [5]int{1, 2, 3, 4, 5}
	for i, v := range myArr {
		pl(i, v)
	}
	a := [4]int{1, 2, 4, 5}
	a[3] = 400
	pl(a)

	arr := [...]int{5: -1, 324, 43}
	pl(arr, "arr")

	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

	// slice 

	slice := make([]int, 6)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	slice[5] = 6
	slice = append(slice, 4, 5, 6, 7, 54, 332, 232)

	for i := 0; i < len(slice); i++ {
		pf("slice[%d] = %d\n", i, slice[i])
	}
	// pl(slice, "slice")
	for _, val := range slice{
		pl(val)
	}
}
