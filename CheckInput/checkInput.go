// This package imports the necessary libraries for reading user input and converting it to an integer

package checkInput

import (
	"fmt"
	"strconv"
)

// This function prompts the user to enter a number and stores it in the variable x
func Check() int {
	var x string
	fmt.Scanln(&x)

	// This code attempts to convert the string x to an integer
	// If the conversion is successful and the integer is between 0 and 3, the value of y is set to the converted integer
	// Otherwise, the value of y is set to 0
	y, err := strconv.Atoi(x)
	if err != nil || y < 0 || y > 3 {
		y = 0
	}
	// This function returns the value of y
	return y
}
