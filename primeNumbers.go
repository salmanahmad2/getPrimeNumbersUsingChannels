package main

import (
	"fmt"
)

func getStartNum(routineNum int, numGo int) int {
	return (routineNum * numGo)
}

func prime(c chan int, numGo int, numStr int) {
	sum := 0
	f := 0
	for i := numStr; i < (numStr + numGo); i++ {
		f = 0
		half := i / 2
		for j := 2; j <= half; j++ {
			if i%j == 0 {
				f++
				break
			}
		}

		if f == 0 && i != 1 && i != 0 {
			sum++
		}
	}

	c <- sum

}

func main() {
	sum := 0
	extra := 0
	minRange := 1
	maxRange := 1000

	if (maxRange-(minRange-1))%5 != 0 {
		extra = (maxRange - (minRange - 1)) % 5
	}

	numGo := (maxRange - (minRange - 1)) / 5
	c := make(chan int)
	go prime(c, numGo, minRange)
	go prime(c, numGo, getStartNum(1, numGo)+minRange)
	go prime(c, numGo, getStartNum(2, numGo)+minRange)
	go prime(c, numGo, getStartNum(3, numGo)+minRange)
	go prime(c, (numGo + extra), getStartNum(4, numGo)+minRange)
	for i := 1; i <= 5; i++ {
		sum = sum + <-c
	}

	fmt.Printf("Total number of prime numbers between the range are:%d", sum)

}
