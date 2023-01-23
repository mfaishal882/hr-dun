package main

func main() {

}

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var count, tempTotal int32
	var mInt int = int(m)
	for i := 0; i < int(len(s))-mInt+1; i++ {
		tempTotal = 0
		for j := i; j < i+mInt; j++ {
			tempTotal += s[j]
		}
		if tempTotal == d {
			count++
		}
	}
	return count
}

/*
test case

Input (stdin)

    1

    4

    4 1

Expected Output

    1

==========

Input (stdin)

    6

    1 1 1 1 1 1

    3 2

Expected Output

    0

	===================

	Input (stdin)

    5

    1 2 1 3 2

    3 2

Expected Output

    2
*/
