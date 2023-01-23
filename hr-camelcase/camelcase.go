package main

import "unicode"

func main() {

}

func camelcase(s string) int32 {
	// Write your code here
	var count int32
	for _, val := range s {
		if unicode.IsUpper(val) {
			count++
		}
	}
	return count + 1
}
