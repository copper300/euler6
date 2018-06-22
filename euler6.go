package main

import (
	"fmt"
)

func main() {
	var sum_sqr int
	var sqr_sum int
	for i := 1; i <= 100; i++ {
		sum_sqr += i
		sqr_sum += i * i
	}
	sum_sqr = sum_sqr * sum_sqr
	fmt.Println(sum_sqr - sqr_sum)
}
