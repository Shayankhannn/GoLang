package main

import "fmt"

// we only have for loop in go
func main() {

// if want to implement while loop 

// this will run in a while fashion

i := 1
 for i <= 3 {
	 fmt.Println("while style for loop",i)
	 i = i + 1
 }

//  for infinite loop 

//  for {
// 	println("Infinite loop")
//  }

// classic for loop

for i := 0; i < 4; i++ {
	
	// you can use contine and break in loop
	// continue will skip the current iteration and move to next iteration
	// break will break the loop

	// if i == 2 {
	// 	continue
	// }
	fmt.Println("classic for loop",i)
}

// we can modernize fro loop with range 

for i:= range 4 {
	fmt.Println("range for loop",i)
}

}

