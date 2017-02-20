package main

import (
	"fmt"
	"math"
	"time"
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


// divide and conquer approach

type TheConcStr struct {
	arr []float64
	u int
	subArr []int
}

func NewConcStr(arr []float64) *TheConcStr {
	tcs := new(TheConcStr)
	tcs.arr = arr
	tcs.u = len(arr)-1
	return tcs
}

func (tcs *TheConcStr) maxSubArr(left int, right int) float64 {
	if left > right {
		return 0.0
	}
	if left == right {
		return tcs.arr[left]
	}
	m := int((left + right) / 2)

	sum := 0.0
	maxToLeft := 0.0
	for i := m; i >= left; i-- {
		sum += tcs.arr[i]
		maxToLeft = math.Max(maxToLeft, sum)
	}

	sum = 0.0
	maxToRight := 0.0
	for i := m + 1; i <= right; i++ {
		sum += tcs.arr[i]
		maxToRight = math.Max(maxToRight, sum)
	}

	maxCrossing := maxToLeft + maxToRight
	maxInL := tcs.maxSubArr(left, m)
	maxInR := tcs.maxSubArr(m+1, right)
	
	return math.Max(math.Max(maxCrossing, maxInL), maxInR)
}

func main() {

	data := []float64{1.0,2.0,-3.0,-4.0,5.0,4.0,-3.0,5.0,-2.0,1.0}

	arr := NewConcStr(data)

	startTime := time.Now()
	fmt.Println(arr.maxSubArr(0,arr.u))
	elapsed := time.Since(startTime)
	fmt.Printf("Execution time: %v\n", elapsed)

}
