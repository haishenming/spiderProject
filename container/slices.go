package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] =", s1)
	s2 := arr[:]
	fmt.Println("arr[:] =", s2)

	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[:2]
	fmt.Println(s2)

	fmt.Println("arr", arr)
	s1 = arr[2:6]
	fmt.Println("s1", s1)
	s2 = s1[3:6]
	fmt.Println("s2", s2)

	// s3, s4, s5 和arr是不一样的数组的切片
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	fmt.Println("arr =", arr)

}
