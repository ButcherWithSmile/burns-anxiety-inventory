package SymptomsOfPhysicalDiscomfort

import (
	t "fmt"
	ci "github.com/Hatef-PR/Burns-Anxiety-Inventory/CheckInput"
	"strings"
)

const options = "0: Not at all | 1: A little | 2: Partially | 3: Very much"

var separator = strings.Repeat("_", 35)

func SymptomsOfPhysicalDiscomfort() int {
	t.Println("Group 3: Symptoms of Physical Discomfort")
	t.Println(separator)

	t.Println("18. Heart beat")
	t.Println(options)
	a := ci.Check()

	t.Println("19. Pain, pressure and tightness in the chest area")
	t.Println(options)
	b := ci.Check()

	t.Println("20. Numbness in the toes or hands")
	t.Println(options)
	c := ci.Check()

	t.Println("21. Anxiety and fear")
	t.Println(options)
	d := ci.Check()

	t.Println("22. Constipation, diarrhea")
	t.Println(options)
	e := ci.Check()

	t.Println("23. Restlessness, jumping")
	t.Println(options)
	f := ci.Check()

	t.Println("24. Muscle cramps")
	t.Println(options)
	g := ci.Check()

	t.Println("25. Sweating without fever")
	t.Println(options)
	h := ci.Check()

	t.Println("26. Feeling of swelling in the throat area")
	t.Println(options)
	i := ci.Check()

	t.Println("27. Trembling and shaking of the body")
	t.Println(options)
	j := ci.Check()

	t.Println("28. Trembling in the leg area")
	t.Println(options)
	k := ci.Check()

	t.Println("29. Feeling dizzy, distracted or losing control")
	t.Println(options)
	l := ci.Check()

	t.Println("30. Suffocation, difficulty breathing")
	t.Println(options)
	m := ci.Check()

	t.Println("31. Pain in the head, neck and back")
	t.Println(options)
	n := ci.Check()

	t.Println("32. Sudden feeling of heat or cold")
	t.Println(options)
	o := ci.Check()

	t.Println("33. Feeling tired, weak and lethargic")
	t.Println(options)
	p := ci.Check()

	t.Println(separator)

	var total = a + b + c + d + e + f + g + h + i + j + k + l + m + n + o + p
	return total
}
