// This package imports the necessary libraries for printing text to the console, getting user input, and converting strings to integers

package worryingThoughts

import (
	t "fmt"
	ci "github.com/Hatef-PR/Burns-Anxiety-Inventory/CheckInput"
	"strings"
)

// This constant stores the options that are presented to the user and the separator that is used to print the options
const options = "0: Not at all | 1: A little | 2: Partially | 3: Very much"

var separator = strings.Repeat("_", 35)

// This function prints the title and separator for the Worrying Thoughts section of the Burns Anxiety Inventory
func WorryingThoughts() int {
	t.Println("Group 2: Worrying thoughts")
	t.Println(separator)

	t.Println("7. Difficulty concentrating")
	t.Println(options)
	a := ci.Check()

	t.Println("8. The mind jumps from one subject to another")
	t.Println(options)
	b := ci.Check()

	t.Println("9. Scary thoughts")
	t.Println(options)
	c := ci.Check()

	t.Println("10. Losing your control")
	t.Println(options)
	d := ci.Check()

	t.Println("11. Fear of madness and emotional disturbance")
	t.Println(options)
	e := ci.Check()

	t.Println("12. Fear of fainting")
	t.Println(options)
	f := ci.Check()

	t.Println("13. Fear of physical illness, heart attack or death")
	t.Println(options)
	g := ci.Check()

	t.Println("14. Fear of being ridiculed")
	t.Println(options)
	h := ci.Check()

	t.Println("15. Fear of being alone and isolated")
	t.Println(options)
	i := ci.Check()

	t.Println("16. Fear of criticism and disapproval")
	t.Println(options)
	j := ci.Check()

	t.Println("17. Prediction of terrible events")
	t.Println(options)
	k := ci.Check()

	t.Println(separator)

	var total = a + b + c + d + e + f + g + h + i + j + k
	return total
}
