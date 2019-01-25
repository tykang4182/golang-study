package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxRun := 500000000
	var probabilityArray [10]int
	var rankingArray [10]int

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
	statistics22Count := 0
	statistics23Count := 0
	statistics24Count := 0
	statistics25Count := 0
	statistics26Count := 0
	statistics27Count := 0
	statistics28Count := 0
	statistics29Count := 0
	statistics30Count := 0
	statistics31Count := 0
	statistics32Count := 0
	statistics33Count := 0
	statistics34Count := 0
	statistics35Count := 0
	statistics36Count := 0
	statistics37Count := 0
	statistics38Count := 0
	statistics39Count := 0
	statistics40Count := 0
	statistics41Count := 0
	statistics42Count := 0
	statistics43Count := 0
	statistics44Count := 0
	statistics45Count := 0

	for times := 1; times <= maxRun; times++ {
		probabilityArray[0] = 1492
		probabilityArray[1] = 483
		probabilityArray[2] = 2054
		probabilityArray[3] = 461
		probabilityArray[4] = 237
		probabilityArray[5] = 976
		probabilityArray[6] = 2578
		probabilityArray[7] = 295
		probabilityArray[8] = 67
		probabilityArray[9] = 1357

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
		} else if (rankingArray[0] == 0 && rankingArray[1] == 7) || (rankingArray[0] == 7 && rankingArray[1] == 0) {
			statistics7Count++
		} else if (rankingArray[0] == 0 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 0) {
			statistics8Count++
		} else if (rankingArray[0] == 0 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 0) {
			statistics9Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 2) || (rankingArray[0] == 2 && rankingArray[1] == 1) {
			statistics10Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 3) || (rankingArray[0] == 3 && rankingArray[1] == 1) {
			statistics11Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 4) || (rankingArray[0] == 4 && rankingArray[1] == 1) {
			statistics12Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 1) {
			statistics13Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 1) {
			statistics14Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 7) || (rankingArray[0] == 7 && rankingArray[1] == 1) {
			statistics15Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 1) {
			statistics16Count++
		} else if (rankingArray[0] == 1 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 1) {
			statistics17Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 3) || (rankingArray[0] == 3 && rankingArray[1] == 2) {
			statistics18Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 4) || (rankingArray[0] == 4 && rankingArray[1] == 2) {
			statistics19Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 2) {
			statistics20Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 2) {
			statistics21Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 7) || (rankingArray[0] == 7 && rankingArray[1] == 2) {
			statistics22Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 2) {
			statistics23Count++
		} else if (rankingArray[0] == 2 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 2) {
			statistics24Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 4) || (rankingArray[0] == 4 && rankingArray[1] == 3) {
			statistics25Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 3) {
			statistics26Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 3) {
			statistics27Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 7) || (rankingArray[0] == 7 && rankingArray[1] == 3) {
			statistics28Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 3) {
			statistics29Count++
		} else if (rankingArray[0] == 3 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 3) {
			statistics30Count++
		} else if (rankingArray[0] == 4 && rankingArray[1] == 5) || (rankingArray[0] == 5 && rankingArray[1] == 4) {
			statistics31Count++
		} else if (rankingArray[0] == 4 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 4) {
			statistics32Count++
		} else if (rankingArray[0] == 4 && rankingArray[1] == 7) || (rankingArray[0] == 7 && rankingArray[1] == 4) {
			statistics33Count++
		} else if (rankingArray[0] == 4 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 4) {
			statistics34Count++
		} else if (rankingArray[0] == 4 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 4) {
			statistics35Count++
		} else if (rankingArray[0] == 5 && rankingArray[1] == 6) || (rankingArray[0] == 6 && rankingArray[1] == 5) {
			statistics36Count++
		} else if (rankingArray[0] == 5 && rankingArray[1] == 7) || (rankingArray[0] == 7 && rankingArray[1] == 5) {
			statistics37Count++
		} else if (rankingArray[0] == 5 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 5) {
			statistics38Count++
		} else if (rankingArray[0] == 5 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 5) {
			statistics39Count++
		} else if (rankingArray[0] == 6 && rankingArray[1] == 7) || (rankingArray[0] == 7 && rankingArray[1] == 6) {
			statistics40Count++
		} else if (rankingArray[0] == 6 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 6) {
			statistics41Count++
		} else if (rankingArray[0] == 6 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 6) {
			statistics42Count++
		} else if (rankingArray[0] == 7 && rankingArray[1] == 8) || (rankingArray[0] == 8 && rankingArray[1] == 7) {
			statistics43Count++
		} else if (rankingArray[0] == 7 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 7) {
			statistics44Count++
		} else if (rankingArray[0] == 8 && rankingArray[1] == 9) || (rankingArray[0] == 9 && rankingArray[1] == 8) {
			statistics45Count++
		}

	}

	fmt.Println("statistics Case 1: ", (float32(statistics1Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 2: ", (float32(statistics2Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 3: ", (float32(statistics3Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 4: ", (float32(statistics4Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 5: ", (float32(statistics5Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 6: ", (float32(statistics6Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 7: ", (float32(statistics7Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 8: ", (float32(statistics8Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 9: ", (float32(statistics9Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 10: ", (float32(statistics10Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 11: ", (float32(statistics11Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 12: ", (float32(statistics12Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 13: ", (float32(statistics13Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 14: ", (float32(statistics14Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 15: ", (float32(statistics15Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 16: ", (float32(statistics16Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 17: ", (float32(statistics17Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 18: ", (float32(statistics18Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 19: ", (float32(statistics19Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 20: ", (float32(statistics20Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 21: ", (float32(statistics21Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 22: ", (float32(statistics22Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 23: ", (float32(statistics23Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 24: ", (float32(statistics24Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 25: ", (float32(statistics25Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 26: ", (float32(statistics26Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 27: ", (float32(statistics27Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 28: ", (float32(statistics28Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 29: ", (float32(statistics29Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 30: ", (float32(statistics30Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 31: ", (float32(statistics31Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 32: ", (float32(statistics32Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 33: ", (float32(statistics33Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 34: ", (float32(statistics34Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 35: ", (float32(statistics35Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 36: ", (float32(statistics36Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 37: ", (float32(statistics37Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 38: ", (float32(statistics38Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 39: ", (float32(statistics39Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 40: ", (float32(statistics40Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 41: ", (float32(statistics41Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 42: ", (float32(statistics42Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 43: ", (float32(statistics43Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 44: ", (float32(statistics44Count) / float32(maxRun) * 100), "%")
	fmt.Println("statistics Case 45: ", (float32(statistics45Count) / float32(maxRun) * 100), "%")

}
