package main

import "fmt"


//? arrays are numbered sequence of specific length


func main() {

	var nums[5]int 
	fmt.Println("lenght",len(nums)) //? built in feature to find length of array
	
//? by default all values are 0 depends on type of array if its int then 0 it just add zeroith value to array

//? zeroith value of types > if int -> 0 , if string -> "" (will look like space) , if bool -> false

fmt.Println("array",nums) 
	
	
	//? we can assign values to array like this
	nums[0] = 1 //? index starts from 0
	nums[1] = 2
	nums[2] = 3
	fmt.Println("array with selective  values",nums[0])

	fmt.Println("array with all values",nums) //? those values which are not assigned are 0

	var string[4]string
	fmt.Println("string array",string) //? default value is empty string
	string[1] = "hello"
	fmt.Println("string array with selective values",string) //? those values which are not assigned are empty string

	var bool[3]bool
	fmt.Println("bool array",bool) //? default value is false

// how to add values to array at the same time of declaring it / to declare it in single line

//? we can do it like this

 number := [3]int{1,2,3} //? if we add extra values it will give error because limit is 3
 fmt.Println("declaring array at the same time adding values",number)

 // 2d array

//?  a 2 x 2 array 

 array2dnums := [2][2]int{{1,2},{3,4}}
fmt.Println("2d array",array2dnums)


}