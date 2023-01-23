package main

import "math"

func main() {}

func catAndMouse(x int32, y int32, z int32) string {
	var result string
	var MouseCatA = math.Abs(float64(z - x))
	var MouseCatB = math.Abs(float64(z - y))

	if MouseCatA == MouseCatB {
		result = "Mouse C"
	} else if MouseCatA < MouseCatB {
		result = "Mouse A"
	} else if MouseCatA > MouseCatB {
		result = "Mouse B"
	}

	return result
}
