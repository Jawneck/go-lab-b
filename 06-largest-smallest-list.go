//A function that returns the largest and smallest elements in a list.
//Author: Danielis Joniskis

package main

//fmt is Go’s standard Input/Output library.
import "fmt"
//sort provides primitives for sorting slices and user-defined collections
import "sort"

func main(){
		//A slice of an integer array.
		x := []int{
			30, 88, -10, 24,
			 6, 76,  12,  2,
			45, 62, -17,  0,
			29, 94,   7, 14,  
		}

		//Sorting a slice of integers in increasing order. https://golang.org/pkg/sort/
		sort.Ints(x)
		lowest := x[0]
		//Displaying the lowest.
		fmt.Println(lowest)
		//Reversing the sort so that we can find the highest integer in the slice.
		sort.Sort(sort.Reverse(sort.IntSlice(x)))
		highest := x[0]
		//Displaying the highest
		fmt.Println(highest)
}