package main

import "fmt"

func main() {}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum, diff int32
	lenBill := int(len(bill))
	for i := 0; i < lenBill; i++ {
		sum += bill[i]
	}

	diff = (sum - bill[k]) / 2

	if b-diff == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - diff)
	}

}
