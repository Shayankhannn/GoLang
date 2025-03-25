package main

import "fmt"

func main() {
	var name string = "John"
	fmt.Println(name)
// infer type
var name2 = "John without staticly declaring type"
fmt.Println(name2)


// bool 
// var isAdult bool = true 
 var isAdult  = true 
fmt.Println(isAdult)

// int 
var age1 int32 = 20
fmt.Println(age1)
// instead of assigning int 32 64 8 you can use int
var age int = 20
fmt.Println(age)

// shorthand no need to write var or type
name3 := "John"
fmt.Println(name3)	

// if you are assigning value later to the variable then you will have to write var and type 

var name4 string
name4 = "John"
fmt.Println(name4)

var totalMarks int
var obtainedMarks int
var percentage float32 

totalMarks = 1100
obtainedMarks = 800
percentage = float32(obtainedMarks)/float32(totalMarks) * 100
fmt.Println(percentage)

result := "Pass"
println(result)

}