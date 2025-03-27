package main

import (
	"fmt"
	"maps"
)

func main()  {
	
// maps are basically associated data structure 
//? maps -> hash -> object -> dictionary

// creating map 
//? for creating map we will use make but with maop keyword
	// syntax for creating map
	//? make(map[key-type]value-type)
	//? map[key-type]value-type

	m := make(map[string]string)
// setting map 
	m["name"] = "golang"
	m["area"] = "backend"

	// printing map 
	//? we print it by its key 

	fmt.Println(m["name"], m["area"]) //? golang
	// fmt.Println(m["area"]) //? backend

	//! Imp : if the key doesnt exist in map it will return zeroith value

	fmt.Println( m["nonexist"]) //? empty string

	//? map2
//? key of string value of int

m2 := make(map[string]int)

m2["1"] = 1
m2["2"] = 2
m2["3"] = 3
m2["4"] = 4

fmt.Println(m2["1"],m2["2"] ) //? 1, 2

//? trying key which doest exist 

fmt.Println(m2["g"]) //? 0 -> zeroith value

// to get length of map

fmt.Println(len(m2)) //? 2 length

// to delete element from map 

//? we have delete func
//? delete(name,key)

delete(m2,"1")

fmt.Println(m2)

// to completely clear out map 
//? we have clear func 
//? clear(name)

clear(m2)
fmt.Println(m2) //? map[]

// short hand / directly declaring maps

m3 := map[int]int{1:1,2:2,3:3,4:4}

fmt.Println(m3)

// to check if value exsist in map 
 //? we can name ok anything but its idiomatic to use ok
 
//? v this value you can name it anything k v --> it contains value 
//? first one return value second one return bool 
 v,ok := m3[2]
 fmt.Println(v) //? 0 --> it returns value .  5 doesnt exist if key exist you will get its value

 if ok {
	fmt.Println("ok")
}else{

	fmt.Println("not ok")
}


// we have a very useful package for map 
//? just like we have slices we have maps with same features

m4 := map[int]int{1:1,2:2,3:3}
m5 := map[int]int{1:1,2:2,3:3}

fmt.Println(maps.Equal(m4,m5))  //? true


}