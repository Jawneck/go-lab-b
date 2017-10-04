//a function that tests whether a string is a palindrome. 
//A palindrome is a string that reads the same in reverse, e.g. Civic.
//Author: Danielis Joniskis

package main

//fmt is Go’s standard Input/Output library.
import "fmt"

//A function testing for a palindrome.
//Taking a string in and returning it as a boolean value.
func palindromeTest(input string) bool {
		for i := 0; i < len(input)/2; i++{
			if input[i] != input[len(input)-i -1]{
				return false
			}
		}
		return true
}

//We test the function in the main method.
func main(){
		fmt.Println(palindromeTest("civic"))
}