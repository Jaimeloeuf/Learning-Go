package main

import "fmt"
import "math"

/*	Go source file for recreating the Square root function in the Math standard library.
	"Sqrt" is the function created to mimic the function.

	@Todos
	- Fix the bugs currently faced.
	- Build in memoization for the Sqrt function
*/

// Shorthand function to calculate square of a value
var SQ = func(input float64) (output float64) { return input * input }

func Sqrt(x float64) float64 {
	fmt.Println("value of X is: ", x)
	z := float64(1)

	for ((z*z - x) / (2 * z)) < -1 {
		z++
		fmt.Println("test: ", (z*z-x)/(2*z))
	}
	return z

	for ; x >= SQ(z+1); z++ {

		if v := SQ(z); x > v {
			continue
		} else {
			return z
		}
	}
	return z
}

func main() {
	x := 67

	// Want to make sure z^2 close to X
	fmt.Println("\nValue from my func: ", Sqrt(float64(x)))
	fmt.Println("Value from Math lib: ", math.Floor(math.Sqrt(float64(x))))
}
