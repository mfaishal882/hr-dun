package main

import "fmt"

func main() {

}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var result int32
	lenA := int(len(a))
	lenB := int(len(b))

	for i := 1; i < 101; i++ {
		for j := 0; j < lenA; j++ {
			if i%int(a[j]) != 0 {
				break
			} else if j == lenA-1 {
				for k := 0; k < lenB; k++ {
					fmt.Println("b[k] = ", b[k], "i =", i)
					if int(b[k])%i != 0 {

						break
					} else if k == lenB-1 {
						result++
					}

				}
			}
		}
	}
	return result
}
