// This package imports the necessary libraries for printing text to the console, getting user input, and converting strings to integers

package worryingFeelings

import (
	t "fmt"
	ci "github.com/Hatef-PR/Burns-Anxiety-Inventory/CheckInput"
	"strings"
)

// This constant stores the options that are presented to the user and the separator that is used to print the options
const options = "0: Not at all | 1: A little | 2: Partially | 3: Very much"

var separator = strings.Repeat("_", 35)

// This function prints the title and separator for the Worrying Feelings section of the Burns Anxiety Inventory
func WorryingFeelings() int {
	t.Println("Group 1: Worrying Feelings")
	t.Println(separator)

	t.Println("1. Anxiety, confusion, worry, fear")
	t.Println(options)
	a := ci.Check()

	t.Println("2. Feeling strange, unreal and vague about things")
	t.Println(options)
	b := ci.Check()

	t.Println("3. A feeling of being separated from all or part of your body")
	t.Println(options)
	c := ci.Check()

	t.Println("4. Sudden and unexpected panic")
	t.Println(options)
	d := ci.Check()

	t.Println("5. A sense of fate")
	t.Println(options)
	e := ci.Check()

	t.Println("6. Cramped feeling, mental pressure")
	t.Println(options)
	f := ci.Check()

	t.Println(separator)

	var total = a + b + c + d + e + f
	return total
}
