package main

import (
	"fmt"
)

func getSubarray(data []int) (int, []int){
	var sum,tmp int
	var c []int
	count := len(data)
	for i,_ := range data {
		if i+1 < count {
			tmp = data[i] + data[i+1]
			if tmp > sum {
				sum = tmp
				c = []int{data[i], data[i+1]}
			}
		}
	}
	return sum, c
}

func main() {
	fmt.Println(getSubarray([]int{1,2,-3,-4,5,4,-3,5,-2,1}))
}
