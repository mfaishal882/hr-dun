package main

import "fmt"

func main() {
	// 6
	// 78
	// 77
	// 76
	// 75
	// 12
	// 39

}

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var result []int32
	var diff int32
	for _, grade := range grades {
		diff = (5 - (grade % 5))
		if grade < 38 {
			result = append(result, grade)
		} else if diff < 3 {
			result = append(result, grade+diff)
			fmt.Println("grade = ", grade, "grade %5 =", (5 - (grade % 5)), "grade+diff", grade+diff)
		} else {
			result = append(result, grade)
		}

	}
	return result
}
