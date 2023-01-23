package main

import "fmt"

func main() {

}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var valueArr []int = []int{1, 2, 3, 4, 5}
	var countArr []int = []int{0, 0, 0, 0, 0}

	for i := 0; i < int(len(arr)); i++ {
		for j := 0; j < len(valueArr); j++ {
			if int(arr[i]) == valueArr[j] {
				countArr[j]++
			}
		}
	}

	for _, val := range countArr {
		fmt.Println(val)
	}

	max := countArr[0]
	var maxIdx int
	for idx, val := range countArr {
		if max < val {
			max = val
			maxIdx = idx
		}
	}

	return int32(valueArr[maxIdx])
}
