//Write a function to reverse a string in go.
//Author: Danielis Joniskis
//Adapted from: https://www.dotnetperls.com/reverse-string-go
package main

//fmt is Go’s standard Input/Output library.
import "fmt"

//We call the reverse function in the main method.
func main(){
	s := "problem sheet"
	fmt.Println(s)

	fmt.Println(reverse(s))
}
//A function which reverses a string.
func reverse(s string) string{
	var reverse string
	for i := len(s) - 1; i >= 0; i--{
		reverse += string (s[i])
	}
	return reverse
}