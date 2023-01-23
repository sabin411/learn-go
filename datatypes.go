package main

import (
	"fmt"
	"reflect"
	"strconv"
)
var pl = fmt.Println

func main(){
	type1 := reflect.TypeOf("sabin");
	pl("The type of the follow string is ",type1)

	type2 := reflect.TypeOf(1234);
	pl("The type of the type2 is ", type2);

	// converting asci to integer
	dec1 := 123.43;
	pl("its has been converted", int(dec1));

	string1 := "1234545";
	convertedInt, convertedIntErr := strconv.Atoi(string1);
	if(convertedIntErr != nil){
		pl("Something went wrong!!!");
	}else{
		pl("Converted integer", convertedInt)
	}


	// converting int to ascii
	int1 := 123456678;
	convertedAscii := strconv.Itoa(int1);

	pl("converted to string", convertedAscii, reflect.TypeOf(convertedAscii), reflect.TypeOf(int1));

}
