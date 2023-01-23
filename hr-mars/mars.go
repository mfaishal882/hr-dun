package main

import "fmt"

func main() {
	fmt.Println(marsExploration("hello"))
}

func marsExploration(s string) int32 {
	// Write your code here

	//var n int = len(s)/3
	var result int32
	var i int = 0
	for i < len(s) {
		fmt.Println("ini " + string(s[i]) + "===" + string(s[i+1]) + "===" + string(s[i+2]) + "\n\n")
		if string(s[i]) != "S" {
			result++
		}
		if string(s[i+1]) != "O" {
			result++

		}
		if string(s[i+2]) != "S" {
			result++

		}
		i += 3
	}

	return result

}
