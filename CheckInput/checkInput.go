package checkInput

import (
	"fmt"
	"strconv"
)

func Check() int {
	var x string
	fmt.Scanln(&x)

	y, err := strconv.Atoi(x)
	if err != nil || y < 0 || y > 3 {
		y = 0
	}
	return y
}
