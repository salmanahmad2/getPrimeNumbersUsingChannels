package main

import (
	"fmt"
)

// prime number between 1-1000
func prime(c chan int, num int) {
	sum := 0
	f := 0

	for i := num; i <= num+199; i++ {
		f = 0
		half := i / 2
		for j := 2; j <= half; j++ {
			if i%j == 0 {
				f++
				break
			}
		}

		if f == 0 && i != 1 {
			sum++
		}
	}
	c <- sum
}

func main() {
	sum := 0
	c := make(chan int)
	go prime(c, 1)
	go prime(c, 201)
	go prime(c, 401)
	go prime(c, 601)
	go prime(c, 801)

	for i := 1; i <= 5; i++ {
		sum = sum + <-c
	}

	fmt.Printf("Total number of prime numbers between 1-1000 are:%d", sum)
}
