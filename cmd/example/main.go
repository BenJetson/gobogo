package main

import (
	"fmt"
	"time"

	bogo "github.com/BenJetson/gobogo"
)

func main() {
	s := []int{2, 7, 5, 8, 4, 3, 10, 0, 6, 9, 1}
	fmt.Println("input:", s)
	begin := time.Now()
	bogo.Sort(s)
	elapsed := time.Since(begin)
	fmt.Println("output:", s)
	fmt.Println("elapsed:", elapsed.Seconds(), "seconds")
}
