//A function to calculate the square root of a number using Newton’s method.
//Author: Danielis Joniskis

package main
//fmt is Go’s standard Input/Output library.
import "fmt"
//math provides basic constants and mathematical functions.
import "math"

//Newtons method is to approximate the square root of a number x by picking a starting point z.
//https://tour.golang.org/flowcontrol/8
func newtonMethod(x float64 ) float64{
	if x == 0 {return 0}
	z := 1.0
	for i := 0; i < int(x); i++ {
		//Repeating the following operation.
		z = z - ((math.Pow(z, 2) -x) /(2 * z))
	}
	return z
}

//The golang Sqrt method
func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

//Here we iterate over a loop to see how close the newtons method is to the sqrt method.
func main(){
	repititions := 10
	for i := 1; i < repititions; i++{
		newtonMethod := newtonMethod(float64 (i))
		Sqrt := Sqrt(float64 (i))
		fmt.Println(i, "squared:")
		fmt.Println( "Newtons Method: ", newtonMethod)
		fmt.Println("  Sqrt:", Sqrt)
	}
}