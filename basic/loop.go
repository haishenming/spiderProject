package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

// 把整数转换为二进制表达式
func convertToBin(n int) string {
	result := ""

	if n == 0{
		return "0"
	}

	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
		convertToBin(83753728),
		convertToBin(0),
	)

	printFile("abc.txt")
}
