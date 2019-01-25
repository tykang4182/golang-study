package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxRun := 1000000
	var probabilityArray [7]int
	var rankingArray [7]int

	statistics1Count := 0
	statistics2Count := 0
	statistics3Count := 0
	statistics4Count := 0
	statistics5Count := 0
	statistics6Count := 0
	statistics7Count := 0
	statistics8Count := 0
	statistics9Count := 0
	statistics10Count := 0
	statistics11Count := 0
	statistics12Count := 0
	statistics13Count := 0
	statistics14Count := 0
	statistics15Count := 0
	statistics16Count := 0
	statistics17Count := 0
	statistics18Count := 0
	statistics19Count := 0
	statistics20Count := 0
	statistics21Count := 0

	for times := 1; times <= maxRun; times++ {
		probabilityArray[0] = 2000
		probabilityArray[1] = 1900
		probabilityArray[2] = 1400
		probabilityArray[3] = 1450
		probabilityArray[4] = 1250
		probabilityArray[5] = 1125
		probabilityArray[6] = 875

		for rank := 0; rank < len(probabilityArray); rank++ {
			var totalNum = 0
			for index := 0; index < len(probabilityArray); index++ {
				totalNum += probabilityArray[index]
			}
			randNum := rand.Intn(totalNum) + 1

			for i := 0; i < len(probabilityArray); i++ {
				randNum -= probabilityArray[i]
				if randNum <= 0 {
					rankingArray[rank] = i
					probabilityArray[i] = 0
					break
				}
			}

		}

		if (rankingArray[0] == 0 && rankingArray[1] == 1) || (rankingArray[0] == 1 && rankingArray[1] == 0) {
			statistics1Count++
		} else if (rankingArray[0] == 0 && rankingArray[1] == 2) || (rankingArray[0] == 2 && rankingArray[1] == 0) {
			statistics2Count++
		} else if (rankingArray[0] == 0 && rankingArray[1] == 3) || (rankingArray[0] == 3 && rankingArray[1] == 0) {
			statistics3Count++
		} else if (rankingArray[0] == 0 && rankingArray[1] == 4) || (rankingArray[0] == 4 && rankingArray[1] == 0) {
			statistics4Count++
		} else if (rankingArray[0] == 0 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 0) {
			statistics5Count++
		} else if (rankingArray[0] == 0 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 0) {
			statistics6Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 2) || (rankingArray[0] == 2 && rankingArray[1] == 1) {
			statistics7Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 3) || (rankingArray[0] == 3 && rankingArray[1] == 1) {
			statistics8Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 4) || (rankingArray[0] == 4 && rankingArray[1] == 1) {
			statistics9Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 1) {
			statistics10Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 1) {
			statistics11Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 3) || (rankingArray[0] == 3 && rankingArray[1] == 2) {
			statistics12Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 4) || (rankingArray[0] == 4 && rankingArray[1] == 2) {
			statistics13Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 2) {
			statistics14Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 2) {
			statistics15Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 4) || (rankingArray[0] == 4 && rankingArray[1] == 3) {
			statistics16Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 3) {
			statistics17Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 3) {
			statistics18Count++
		} else if (rankingArray[0] == 4 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 4) {
			statistics19Count++
		} else if (rankingArray[0] == 4 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 4) {
			statistics20Count++
		} else if (rankingArray[0] == 5 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 5) {
			statistics21Count++
		}

	}
	fmt.Println("statistics Case 1: ", (float32(statistics1Count) / float32(maxRun)))
	fmt.Println("statistics Case 2: ", (float32(statistics2Count) / float32(maxRun)))
	fmt.Println("statistics Case 3: ", (float32(statistics3Count) / float32(maxRun)))
	fmt.Println("statistics Case 4: ", (float32(statistics4Count) / float32(maxRun)))
	fmt.Println("statistics Case 5: ", (float32(statistics5Count) / float32(maxRun)))
	fmt.Println("statistics Case 6: ", (float32(statistics6Count) / float32(maxRun)))
	fmt.Println("statistics Case 7: ", (float32(statistics7Count) / float32(maxRun)))
	fmt.Println("statistics Case 8: ", (float32(statistics8Count) / float32(maxRun)))
	fmt.Println("statistics Case 9: ", (float32(statistics9Count) / float32(maxRun)))
	fmt.Println("statistics Case 10: ", (float32(statistics10Count) / float32(maxRun)))
	fmt.Println("statistics Case 11: ", (float32(statistics11Count) / float32(maxRun)))
	fmt.Println("statistics Case 12: ", (float32(statistics12Count) / float32(maxRun)))
	fmt.Println("statistics Case 13: ", (float32(statistics13Count) / float32(maxRun)))
	fmt.Println("statistics Case 14: ", (float32(statistics14Count) / float32(maxRun)))
	fmt.Println("statistics Case 15: ", (float32(statistics15Count) / float32(maxRun)))
	fmt.Println("statistics Case 16: ", (float32(statistics16Count) / float32(maxRun)))
	fmt.Println("statistics Case 17: ", (float32(statistics17Count) / float32(maxRun)))
	fmt.Println("statistics Case 18: ", (float32(statistics18Count) / float32(maxRun)))
	fmt.Println("statistics Case 19: ", (float32(statistics19Count) / float32(maxRun)))
	fmt.Println("statistics Case 20: ", (float32(statistics20Count) / float32(maxRun)))
	fmt.Println("statistics Case 21: ", (float32(statistics21Count) / float32(maxRun)))

}
