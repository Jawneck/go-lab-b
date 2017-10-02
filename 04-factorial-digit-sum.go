// A program that calculates the sum of the digits of the factorial of a number.
//Author: Danielis Joniskis

package main
//fmt is Go’s standard Input/Output library.
func P20() int {
    sum := 0;
    digits := [200]int{};
    digits[0] = 1;
    for i := 2; i <= 100; i++ {
    	for j := 0; j < len(digits); j++ {
    		digits[j] *= i;
    		if j > 0 && digits[j - 1] > 9 {
    			digits[j] += int(digits[j - 1] / 10);
    			digits[j - 1] %= 10;
    		}
    	}
    }
    for i := 0; i < len(digits); i++ {
    	sum += digits[i];
    }
    return sum;
}


