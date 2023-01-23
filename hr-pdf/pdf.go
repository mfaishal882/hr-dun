package mamin

import "fmt"

func main() {
	fmt.Println("lalala")
}

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here

	var maxHeight int32 = 0
	for idx, _ := range word {
		for idx2, val2 := range h {
			if []rune(word)[idx]-97 == int32(idx2) {
				if maxHeight < val2 {
					maxHeight = val2
				}
			}
		}
	}

	return int32(len(word)) * maxHeight
}
