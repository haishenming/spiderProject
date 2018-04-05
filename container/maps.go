package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "ccmouse",
		"caurse": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	//caurseName, ok := m["caurse"]
	//fmt.Println(ok, caurseName)

	// 常用写法, ok 是是否能拿到值
	if caurseName, ok := m["caurse"]; ok {
		fmt.Println(caurseName)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println("Deleting values")
	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)

}
