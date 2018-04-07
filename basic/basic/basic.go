package basic

import (
	"fmt"
	"math/cmplx"
	"math"
)

var aa = 3
var ss = "kkk"
var bb = true

func variableZeroValue() {
	var a int
	var s string

	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, s = 3, 4, true, "def"

	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5

	fmt.Println(a, b, c, s)
}

func euler()  {
	// 欧拉公式
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

func tringle() {
	var a, b int = 3, 4
	var c int

	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)

}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 常量
func consts() {
	const filename string = "abc.txt"
	const a, b =3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举
func enums() {
	const(
		cpp = iota
		java
		python
		golang
	)

	const(
		b = 1 << (10 * iota)
		kb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, gb, tb, pb)
}

func bounded(v int) int {
	if v < 100 {
		return 100
	} else if v < 10 {
		return 0
	} else {
		return v
	}
}

func main() {
	fmt.Println("Hello world")

	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, ss, bb)

	euler()
	tringle()

	// 常数
	const (
		con  = 1 << 1
		n1  float32 = iota * con
		n2  float32 = iota * con
		n3  float32 = iota * con
		n4  float32 = iota * con
	)

	fmt.Println(con, n1, n2, n3, n4)

	consts()

	enums()
}
