package main

func main() {

}

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	onTime := 0
	late := 0
	for _, val := range a {
		if val <= 0 {
			onTime++
		} else {
			late++
		}
	}

	if k > int32(onTime) {
		return "YES"
	} else {
		return "NO"
	}
}
