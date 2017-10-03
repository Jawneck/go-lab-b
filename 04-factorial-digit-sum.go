// A program that calculates the sum of the digits of the factorial of a number.
//Author: Danielis Joniskis

package main
//fmt is Go’s standard Input/Output library.
import "fmt"
//strings implements simple functions to manipulate UTF-8 encoded strings.
import "strings"
//strconv implememts conversions to and from strings.
import "strconv"
import "math/big"
//A function to find a factorial of n. 'uint64' is the set of all unsigned 64-bit integers.
func factorial(n uint64)(result uint64){
		if (n > 0){
				result = n * factorial(n-1)
				return result
		}
		return 1
}		

//A function to sum all the individual integers of the factorial n.
func sumDigits(result string) int {
			  //slices string into all substrings.
	digits := strings.Split(result, "")

	total := 0
	// _ avoids having to declare all variables for return type. It is called the blank identifier.
	//The Go compiler won't let you create variables that you never use.
	// range iterates over elements, we are using it to sum the digits.
	for _, digit := range digits{
				  //strconv.Atoi converts string to int.
		val, _ := strconv.Atoi(digit)
		total += val
	}
	return total
}

func main(){
	
	var n uint64

	n = factorial(100)
	fmt.Println(n)
			//Converting the value to a string
	var x = strconv.FormatUint(uint64(n),100)

	fmt.Println(sumDigits(x))
}
