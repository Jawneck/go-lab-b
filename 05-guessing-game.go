//A program where the user must guess a randomly generated number.
//After every guess the program tells the user whether their number was too high or too low.
//At the end, the number of tries should be printed. 
//It counts only as one try if they input the same number multiple times consecutively.
//Author: Danielis Joniskis

package main

//fmt is Go’s standard Input/Output library.
import "fmt"
//rand implements psuedo-random number generation 
import "math/rand"
//time provides functionality for measuring and displaying time.
import "time"




func main() {
	//Giving a random see so that the same value is not generated every time.
	rand.Seed(time.Now().UnixNano())
	//Generating a random int
	randomNum := rand.Intn(100)

	//Setting the counter.
	i := 1

	for {
		var guess int
		fmt.Println("Guess a number between 1 and 100: ")
		fmt.Scanln(&guess)

		 if guess == randomNum{
			fmt.Println("You have guessed the random number!")
			fmt.Printf("\nThe number of times you have guessed is: %d", i)
			return
		} else if guess < randomNum{
			fmt.Println("Guess higher!")
		} else if guess > randomNum{
			fmt.Println("Guess lower!")
		}else{
			fmt.Println("Please enter an integer value.")
		}

		i++

	}

}