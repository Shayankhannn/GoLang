package main

import "fmt"


//? to accept n number of parameter in sum we can do (parameter ...type) 

//? now nums is like we get  slice  so what wee can do is loop
func sum(nums ...int) int  {
	total := 0
	for _,num := range nums {
		total = total + num
	}
	return total
}

func main()  {
	
// variadic func are those in which you can pass n number of parameter 
//? like Println(...)

result := sum(2,3,4,5,6)
fmt.Println(result)

}