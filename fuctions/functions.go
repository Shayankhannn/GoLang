package main

import "fmt"

//? if both parameter are of dame type you can write like this too (a,b int)
func add(a int, b int) int  {
	return a + b
}

// go func can return multiple values

//? if multiple return you write in bracket () and mention return type

//! the return type should be in same order
func getLanguages()(string,bool,string)  {
	return "golang", true, "C"
}

// in go functions are first class citizen functions

//? which means we can assign afunc to a variable too or pass it in other func as a argument too
//? or from one function return a function too

func processIt(fn func(a int)int)  {
	fn(3)
}

//? or if you want to return a func

func processIt2() func(a int)int{
 return func(a int)int{
	return a
 }
}


func main()  {
	
	fmt.Println(add(10,10))
	fmt.Println(getLanguages())
	//? you can do it llike this too

	lang1,lang2, _ := getLanguages()

	fmt.Println(lang1,lang2)
//? sometimes we dont need to get all vluse like 
//? we dont want to print lang 2 so tu suppress it we can write _

fn := func (a int) int{
return 2
}
processIt(fn)

fn2 := processIt2()

fmt.Println(
	fn2(3))

}