package main

import "fmt"

func LengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i

	}
	return maxLength

}

func main() {
	fmt.Println(
		LengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(LengthOfNonRepeatingSubStr("红鲤鱼与绿鲤鱼与驴"))

}
