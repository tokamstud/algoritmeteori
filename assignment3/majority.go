package main

import (
	"fmt"
)

func hasMajorityGreedy(data []int) bool {
	count := 0
	n := len(data)
	for i, _ := range data {
		//me := data[i]
		for j, _ := range data {
			if j < len(data) && i < len(data){
				if data[i] != data[j] {
					data = append(data[:j], data[j+1:]...)
					data = append(data[:j+1], data[j+2:]...)
					count += j
					break
				}
			}
		}
	}
	if count >= n {
		return true
	}
	return false
}

func hasMajority(data []int) bool {
	var me int
	count := 0
	for i := range data {
		if (count == 0) {
			me = data[i]
		}
		if me == data[i] {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, e := range data {
		if e == me {
			count++
		}
	}
	if count > len(data)/2 {
		return true
	}
	return false
}

func main() {

	data1 := []int{2, 3, 3, 2, 3, 3}
	data2 := []int{6, 3, 2, 7, 3, 1}
	data3 := []int{3,3,3,2,2,2,1,3,2,3,3,3}

	fmt.Println(hasMajority(data1))
	fmt.Println(hasMajority(data2))
	fmt.Println(hasMajority(data3))

	fmt.Println(hasMajorityGreedy(data1))
	fmt.Println(hasMajorityGreedy(data2))
	fmt.Println(hasMajorityGreedy(data3))

}
