//A function that merges two sorted lists into a new sorted list,
//e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].
//Author: Danielis Joniskis

package main

//fmt is Go’s standard Input/Output library.
import "fmt"
//sort provides primitives for sorting slices and user-defined collections
import "sort"

func main(){

	//Concatenating the two slices. https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
	slice := append([]int {1, 4, 6}, []int {2, 3, 5}...)

	fmt.Println("The two integer lists concatenated:", slice)
	//Sorting the concatenated slice
	sort.Ints(slice)
	fmt.Println("The new sorted list:", slice)
}

