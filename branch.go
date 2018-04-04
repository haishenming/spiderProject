package main

import (
	"io/ioutil"
	"fmt"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a + b
	case "/":
		result = a / b
	default:
		panic("Unsupported operator: " + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))

	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	result := eval(1, 2, "+")
	fmt.Println(result)

	fmt.Println(
		grade(0),
		grade(1),
		grade(-1),
		grade(100),
		grade(79),
	)

}
