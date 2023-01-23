package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	sort.Slice(fruits, func(i, j int) bool {
		return fruits[i] < fruits[j]
	})

	fmt.Println(fruits)
}

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	sockSlice := ar[0:len(ar)]
	var pair int

	sort.Slice(sockSlice, func(i, j int) bool {
		return sockSlice[i] < sockSlice[j]
	})

	for i := 0; i < int(n-1); i++ {
		if ar[i] == ar[i+1] {
			pair++
			i++
		}
	}

	fmt.Println(sockSlice)
	fmt.Println(pair)
	return int32(pair)
}
