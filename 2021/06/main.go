package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	days  = 80
	days2 = 256
)

func iteration(num []int64) []int64 {
	next := make([]int64, 9)
	next[6] = num[0]
	next[8] = num[0]
	for j := 1; j < len(num); j++ {
		next[j-1] += num[j]
	}
	return next
}

func count(num []int64) int64 {
	var n int64
	for _, nn := range num {
		n += nn
	}
	return n
}

func main() {
	num := make([]int64, 9) // 0-8
	for i, s := range strings.Split(input, ",") {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil || int(n) >= len(num) {
			// should not happen
			log.Printf("input #%d malformed: %v (%q)", i, err, s)
			continue
		}
		num[int(n)]++
	}
	for i := 0; i < days; i++ {
		num = iteration(num)
	}
	fmt.Println(count(num))
	for i := days; i < days2; i++ {
		num = iteration(num)
	}
	fmt.Println(count(num))
}

const input = `3,1,4,2,1,1,1,1,1,1,1,4,1,4,1,2,1,1,2,1,3,4,5,1,1,4,1,3,3,1,1,1,1,3,3,1,3,3,1,5,5,1,1,3,1,1,2,1,1,1,3,1,4,3,2,1,4,3,3,1,1,1,1,5,1,4,1,1,1,4,1,4,4,1,5,1,1,4,5,1,1,2,1,1,1,4,1,2,1,1,1,1,1,1,5,1,3,1,1,4,4,1,1,5,1,2,1,1,1,1,5,1,3,1,1,1,2,2,1,4,1,3,1,4,1,2,1,1,1,1,1,3,2,5,4,4,1,3,2,1,4,1,3,1,1,1,2,1,1,5,1,2,1,1,1,2,1,4,3,1,1,1,4,1,1,1,1,1,2,2,1,1,5,1,1,3,1,2,5,5,1,4,1,1,1,1,1,2,1,1,1,1,4,5,1,1,1,1,1,1,1,1,1,3,4,4,1,1,4,1,3,4,1,5,4,2,5,1,2,1,1,1,1,1,1,4,3,2,1,1,3,2,5,2,5,5,1,3,1,2,1,1,1,1,1,1,1,1,1,3,1,1,1,3,1,4,1,4,2,1,3,4,1,1,1,2,3,1,1,1,4,1,2,5,1,2,1,5,1,1,2,1,2,1,1,1,1,4,3,4,1,5,5,4,1,1,5,2,1,3`
