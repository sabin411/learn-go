package main

import (
	"fmt"
)

var pl = fmt.Println;
type DataProps struct {
	name string;
	username string;
	password string;
	age int;
}

type UserProps struct {
	school string;
	userInfo DataProps;
}

func nameChanger(dataObj *DataProps){
		dataObj.name = "Saraswoti pandey";
}
func main(){
	// Create a new instance of the struct
	userInfo  := &DataProps{name: "sabin"}
	pl(userInfo, userInfo.name)
	nameChanger(userInfo);
	pl(userInfo, userInfo.name)

	info := UserProps{"Lincoln internation school", DataProps{"sabin", "sabin411", "password", 12}}

	pl(info, "info");
}

