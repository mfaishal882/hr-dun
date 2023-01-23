package main

func main() {}

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var max, result int32

	for i := 0; i < int(len(height)); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	if k < max {
		result = max - k
	}

	return result
}
