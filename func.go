package main

import "fmt"

func div(a, b int) (q, r int) {
	q, r = a / b, a % b
	return
}


func main() {
	fmt.Println(div(13, 3))
}
