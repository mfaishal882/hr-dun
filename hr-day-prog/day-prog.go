package main

import "fmt"

func main() {

}

func dayOfProgrammer(year int32) string {
	// Write your code here
	// var dayOfProgrammer string = fmt.Sprint("13.09.", year)

	// if (year%4 == 0 && year%100 > 0) || year%400 == 0 {
	// 	dayOfProgrammer = fmt.Sprint("12.09.", year)
	// } else if (year%4 == 0 && year%100 == 0) || year < 2000 {
	// 	dayOfProgrammer = fmt.Sprint("12.09.", year)
	// } else if year == 1918 {
	// 	dayOfProgrammer = fmt.Sprint("26.09.", year)
	// }

	var result string
	if year > 1918 {
		if (year%4 == 0 && year%100 > 0) || year%400 == 0 {
			result = fmt.Sprint("12.09.", year)
		} else {
			result = fmt.Sprint("13.09.", year)
		}
	} else if year < 1918 {
		if year%4 == 0 {
			result = fmt.Sprint("12.09.", year)
		} else {
			result = fmt.Sprint("13.09.", year)
		}
	} else {
		result = "26.09.1918"
	}

	return result
}
