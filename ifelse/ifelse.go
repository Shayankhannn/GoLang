package main

import "fmt"

func main(){
	age := 10

	if age >= 20 {
		fmt.Println("person is abult")
	}else {
		fmt.Println("person is not abult")
	}

	// if else if 

	age1 := 12

	if age1 >= 20 {
		fmt.Println("person is abult")
	}else if age1 >= 10 {
		fmt.Println("person is teenager")
	}else {
		fmt.Println("person is child")
	}

	// logical operators with if else 

	var role string = "admdin"
	user := false

	if role == "admin " || user {
		fmt.Println("Access granted")
	}else {
		fmt.Println("Access denied")
	}

	// we can declare a variable in if else statement 

	// and the good thing is that we can use it inside the if else block 
	
	//! you cant use that variable outside the if else block

	//? go doesnt have ternary operator you will have to use if else


	if age2 := 18; age2 >= 18 {
		fmt.Println("person is abult",age2)
	}


}