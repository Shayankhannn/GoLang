package main

import (
	"fmt"
	"slices"
)

//? no need to declare lenght of array in slices
//? most used construct in go
func main() {

//? uninitialized slice is nil
var slice[]int
fmt.Println("uninitialized slice",slice) //? nil

//? you can check if its nill to fiind out
fmt.Println("uninitialized slice",slice == nil) //? true

//? checking length of slice
fmt.Println("lenght of slice",len(slice)) //? 0

//? creating slice with make function
//? make isa built in function to create slices
//? make([]type, length, capacity)

slice2 := make([]int,3) //? 3 is length of slice 

fmt.Println("slice2-makefunc",slice2) //? [0 0 0]

//? to check capacity of slice
fmt.Println("capacity of slice2",cap(slice2)) //? 3

//? to add values to make slice there is append function
//? append(slice, values)

slice2append := append(slice2,1,2,3,4,5,6)
fmt.Println("slice2-append",slice2append) //? [1 2 3]

//? to check capacity of slice
fmt.Println("capacity of slice2append",cap(slice2append)) //? 6


// when you reach the capacity of slice it will double the capacity of slice and double the size of slice cuz its dynamic

//? as you can see that first three values of slice 2 are zero because of make because we gave initial length of 3

slice3 := make([]int,0,3) //? 3 is capacity of slice and we keep zero for initial so that the vlues of zero does not appear
fmt.Println("slice3-makefunc",slice3) 

//? now if you append values to it it will add values to it but there will be no initial values of zero in it

slice3append := append(slice3,1,2,3,4,5,6)
fmt.Println("slice3-append",slice3append) //? [1 2 3 4 5 6]
fmt.Println("capacity of slice3append",cap(slice3append)) //? 6
fmt.Println("lenght of slice3append",len(slice3append)) //? 6

//? if you want to add values to zeroith value in make slice you can do that by using index
//? slice[index] = value
//? but ou have to make sure the initial value is not zero in make slice
//? 3 -> initial length of slice 4 -> capacity of slice 
zeroithappend := make([]int,3,4)
fmt.Println("zeroithappend",zeroithappend) //? [0 0 0]

zeroithappend[0] = 1
fmt.Println("zeroithappend-but index added value",zeroithappend) //? [1 0 0]


// the other way to create slice is to use slice literal


//? slice literal is like array literal but without length
//? []type{values}

sliceliteral := []int{}
fmt.Println("sliceliteral",sliceliteral) //? []
sliceliteral = append(sliceliteral,1,2,3,4,5,6)
fmt.Println("sliceliteral-with value appended",sliceliteral) //? [1 2 3 4 5 6]
fmt.Println("sliceliteral-cap",cap(sliceliteral)) //? 0
fmt.Println("sliceliteral-len",len(sliceliteral)) //? 0


// copy function is used to copy values from one slice to another
//? go has built in copy function

//? copy(destination, source)

var copyslice = make([]int,0,5)

copyslice = append(copyslice,2)
var copyslice2 = make([]int,len(copyslice))

fmt.Println("copyslice",copyslice) //? [2]
fmt.Println("copyslice2",copyslice2) //? [0]

//? copy function
copy(copyslice2,copyslice)
fmt.Println("copyslice2-after copy",copyslice2) 

//? copyslice2 is still empty due to length of copyslice2 is zero if the destination length is zero nothing will be copy to it
//? so we just move copyslice append above copyslice 2 in that way copyslice2 will have values in it because its len is connected to copyslice 

//? now copy func will work

//? both slices after copying
fmt.Println("copyslice",copyslice,"copyslice2",copyslice2) //? [2] [2]



// slice operator
//? slice operator is used to get values from slice
//? slice[low:high] / slice[fromIndex:toIndex] -- if we skip from index it will start from 0 and if we skip to index it will go to end of slice [low:] and [:high]
//? low is starting index and high is ending index but it exclude the high index value

var sliceoperator = []int{1,2,3,4}
fmt.Println("sliceoperator",sliceoperator[0:2]) //? [1 2 ]
fmt.Println("sliceoperator",sliceoperator[:1]) //? [1]
fmt.Println("sliceoperator",sliceoperator[1:]) //? [2,3,4]

// slices package

var slicepackage = []int{1,2,3,4,5,6,7,8,9,10}
var slicepackage2 = []int{1,2,3,4,5,6,7,8,9,10}

fmt.Println("slicepackage", slices.Equal(slicepackage,slicepackage2)) //? true
                        // slices have many functions like equal,sort,reverse etc


// 2d slices
//? just like array we can create 2d arrays as well 

var slice2d = [][]int{{1,2,3},{4,5,6}}

fmt.Println("2d slice",slice2d)

}
