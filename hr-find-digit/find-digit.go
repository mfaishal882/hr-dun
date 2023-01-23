package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func findDigits(n int32) int32 {
	// Write your code here
	var result int32
	var valInt int
	for _, val := range strconv.Itoa(int(n)) {
		fmt.Println(string(val))
		valInt, _ = strconv.Atoi(string(val))
		fmt.Println(valInt)
		if valInt == 0 {
			continue
		}
		if int(n)%valInt == 0 {
			result++
		}
	}
	return result
}
