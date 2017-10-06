//A program that prints the current time and date to the console.
//Author: Danielis Joniskis
//Adapted from https://tour.golang.org/welcome/4
package main

//Imports
//fmt is Go’s standard Input/Output library.
import "fmt"
//time provides functionality for measuring and displaying time.
import "time"

func main(){
	//The Time returned by time.Now contains a monotonic clock reading.
	fmt.Println("The current time and date is ", time.Now())
}//end main