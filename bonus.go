package main

import "fmt"
import "math/rand"

func main() {
	maxRun := 15
	var BonusprobabilityArray [21]int
	var BonusrankingArray [21]int

	for times := 1; times <= maxRun; times++ {
		BonusprobabilityArray[0] = 1
		BonusprobabilityArray[1] = 11
		BonusprobabilityArray[2] = 10
		BonusprobabilityArray[3] = 10
		BonusprobabilityArray[4] = 20
		BonusprobabilityArray[5] = 29
		BonusprobabilityArray[6] = 5
		BonusprobabilityArray[7] = 5
		BonusprobabilityArray[8] = 8
		BonusprobabilityArray[9] = 15
		BonusprobabilityArray[10] = 20
		BonusprobabilityArray[11] = 5
		BonusprobabilityArray[12] = 8
		BonusprobabilityArray[13] = 8
		BonusprobabilityArray[14] = 20
		BonusprobabilityArray[15] = 2
		BonusprobabilityArray[16] = 10
		BonusprobabilityArray[17] = 10
		BonusprobabilityArray[18] = 1
		BonusprobabilityArray[19] = 1
		BonusprobabilityArray[20] = 1

		for rank := 0; rank < len(BonusprobabilityArray); rank++ {
			var totalNum = 0
			for index := 0; index < len(BonusprobabilityArray); index++ {
				totalNum += BonusprobabilityArray[index]
			}
			randNum := rand.Intn(totalNum) + 1

			for i := 0; i < len(BonusprobabilityArray); i++ {
				randNum -= BonusprobabilityArray[i]
				if randNum <= 0 {
					BonusrankingArray[rank] = i
					BonusprobabilityArray[i] = 0
					break
				}
			}

		}
		fmt.Println("times: ", times, " => ", BonusrankingArray[0], BonusrankingArray[1])
	}

}
