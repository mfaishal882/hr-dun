package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var min, max, countMin, countMax int32
	var result []int32
	for idx, score := range scores {
		fmt.Println("Before", score, "min = ", min, "countMin = ", countMin, "max =", max, "countMax =", countMax)
		if idx == 0 {
			min = score
			max = score
		} else {
			if score > min {
				countMin++
				min = score
			}
			if score < max {
				countMax++
				max = score
			}
		}
		fmt.Println(idx+1, " After score = ", score, "min = ", min, "countMin = ", countMin, "max =", max, "countMax =", countMax, "\n\n==================")
	}
	result = append(result, countMin)
	result = append(result, countMax)
	return result
}

/*
Input (stdin)
    9
    10 5 20 20 4 5 2 25 1
Expected Output
    2 4

Input (stdin)
    10
    3 4 21 36 10 28 35 5 24 42
Expected Output
    4 0
*/

func main() {

}
