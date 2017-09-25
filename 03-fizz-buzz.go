//A program that prints the numbers from 1 to 100, each on a new line, to the console, except for the following conditions. 
//For multiples of three print Fizz instead of the number, and for multiples of five print Buzz. 
//For numbers that are multiples of both three and five print FizzBuzz.
//Author: Danielis Joniskis
//Adapted from http://wiki.c2.com/?FizzBuzzTest
package main

//fmt is Go’s standard Input/Output library.
import "fmt"

func main() {
	//Opening the for loop
	for i := 0; i <= 100; i++{
		//Checking if number i is divisable by 15, if it is we dont check the other two conditions.
		if(i % 15 == 0){
			fmt.Println("Fizz" + "Buzz")
		} else if (i % 3  == 0) {
			fmt.Println ("Fizz")
		} else if (i % 5 == 0){
			fmt.Println ("Buzz")
		} else{
			fmt.Println(i)
		}		
	}
}//end main