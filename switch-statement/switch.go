package main

import (
	"fmt"
	"time"
)

func main(){

	//! simple switch

	//? switch expression {
	//? case condition:
	//? }

	i := 1 

	switch i {
	case 1 :
		fmt.Println("one")
		// you dont have to write break keyword go automatically break it
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		//? default is optional

		fmt.Println("other")
	}

	// multiple condition switch

	//? time in standard in built library


	switch time.Now().Weekday() {

//? we gave multiple condition in one case
	
case time.Saturday, time.Sunday :
		fmt.Println("weekend")
		default : 
		fmt.Println("weekday")
	}

// type switch

//? we can save func in variable

whoAmI := func(i interface{}) {
	//! if you dont want to use this value t :=  you can directly do it to like this:
	//! switch i.(type){

	switch t := i.(type){
	case int :
		fmt.Println("I am int",t)
	case string :
		fmt.Println("I am string",t)
	case bool : 
		fmt.Println("I am bool",t)
	default:
		fmt.Println("I am unknown",t)
	}
}

whoAmI(1)
whoAmI("hello")
whoAmI(true)
whoAmI(1.0)

}