package main

import "math"

func main() {

}

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var iReverse int32 = 0
	var counter int32 = 0
	for i <= j {
		iReverse = reverseNumber(i)
		if int32(math.Abs(float64(i-iReverse)))%k == 0 {
			counter++
		}
		i++
	}
	return counter
}

func reverseNumber(num int32) int32 {

	var res int32 = 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
