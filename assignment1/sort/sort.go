package main

import (
	"fmt"
	"math/rand"
	"time"
)
const BIG = 100000
const SMALL = 10
// Insertion sort struct
type Insertion struct {
	data []int
}

func NewInsertion(data []int) *Insertion {
	in := &Insertion{}
	in.data = data
	return in
}

func (in *Insertion) Sort() []int {
	for i := 1; i < len(in.data); i++ {
		tmp := in.data[i]
		hole := i
		for j := i; j > 0 && in.data[j-1] > tmp; j-- {
			in.data[j] = in.data[j-1]
			hole--
		}
		in.data[hole] = tmp
	}
	return in.data
}

// Merge sort struct
type Merge struct {
	data []int
	temp []int
	count int
}

func NewMerge(data []int) *Merge {
	me := &Merge{}
	me.data = data
	me.count = len(data)
	return me
}

func (me *Merge) Sort(data []int) []int {
	if len(data) == 1 {
		return data
	}
	count := len(data)

	var left, right []int
	for i := 0; i < count/2; i++ {
		left = append(left, data[i])
	}
	for i := count/2; i < count; i++ {
		right = append(right, data[i])
	}

	left = me.Sort(left)
	right = me.Sort(right)

	return me.merge(left, right)
}
func (me *Merge) merge(left []int, right []int) []int {
	var c []int
	for len(left)>0 && len(right)>0 {
		if left[0] > right[0] {
			c = append(c,[]int{right[0]}...)
			right = append(right[:0], right[1:]...)
		} else {
			c = append(c,[]int{left[0]}...)
			left = append(left[:0], left[1:]...)
		}
	}

	for len(left)>0 {
		c = append(c,[]int{left[0]}...)
		left = append(left[:0], left[1:]...)
	}

	for len(right)>0 {
		c = append(c,[]int{right[0]}...)
		right = append(right[:0], right[1:]...)
	}
	return c
}

// unexported functions

func genData(size int) []int {
	// generate data:
	// 0 gives small set
	// 1 gives huge set
	var data []int
	rand.Seed(42)
	if size==1 {
		for i := 0; i < BIG; i++ {
			data = append(data, rand.Intn(BIG))
		}
	} else {
		for i := 0; i < SMALL; i++ {
			data = append(data, rand.Intn(SMALL))
		}
	}
	return data
}

func main() {

	big := genData(1)
	small := genData(0)

	in := NewInsertion(small)
	start := time.Now()
	in.Sort()
	elapsed := time.Since(start)
	fmt.Printf("Insertion sort with SMALL data took %s\n", elapsed)

	me := NewMerge(small)
	start = time.Now()
	me.Sort(me.data)
	elapsed = time.Since(start)
	fmt.Printf("Merge sort with SMALL data took %s\n", elapsed)


	inbig := NewInsertion(big)
	start = time.Now()
	inbig.Sort()
	elapsed = time.Since(start)
	fmt.Printf("Insertion sort with BIG data took %s\n", elapsed)

	mebig := NewMerge(big)
	start = time.Now()
	mebig.Sort(mebig.data)
	elapsed = time.Since(start)
	fmt.Printf("Merge sort with BIG data took %s\n", elapsed)


}
