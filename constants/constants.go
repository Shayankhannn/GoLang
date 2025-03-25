package main

import "fmt"

// you can declare variables outside function too 

const pi = 3.14

// you cant declare variable outside function like this
// name := "John"

func main() {
	
const  name string = "John"

// you cant reassign the value of the constant
// name = "Doe"

const age int = 20
fmt.Println(age)


fmt.Println(name)

// you can use variable declare outside function iside function
fmt.Println(pi)


// you can group multiple constants together called grouping 

const (
	port = 8080
	host = "localhost"
)

// fmt.Println(port,host )

fmt.Println(port)
fmt.Println(host)



}