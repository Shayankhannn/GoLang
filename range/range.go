package main

import "fmt"

func main()  {
	// range is used for iterating over data structure
	
	nums := []int{1,2,3,4,5}
	//? if we want to iterate over this slice either we can use for loop

	//? first for loop

	for i :=0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

//?  using range

for i:= range nums  {
	fmt.Println("range",i ) //? excluded the last value
}

//? suming using range

sum := 0

for _, num := range nums {
	sum = sum + num
}

fmt.Println(sum)

// now for map

m := map[int]string{1:"apple",2:"mango",3:"pine"}

//? single value like only k will return only key

for k,v := range m {
	fmt.Println(k,v)
} 

// we can use range on string too

for k,v := range "apple" {
	//! rune is data type in go
	//? its like if a alphabet is under 255 it take 1 byte and if its more than 255 it take 2 byte 
	//? so index then wouldnt be 0 1 bec first alphabet took the 0 1 so the next one will have 2 
	//? this is what starting byte is
	 //? v will give  unicode code point for alphabet rune
	//? and k will basically give starting byte of rune
	//? so if you want each character you can use string func		
	fmt.Println(k,string(v)) //? a p p l e

}





}

