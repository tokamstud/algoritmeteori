package main

import (
	"container/heap"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
)

type Node struct {
	left   *Node
	right  *Node
	freq   int
	val string
}

func NewNode(val string, freq int) *Node {
	return &Node{
		val: val,
		freq: freq,
	}
}

type NodeHeap []Node

func (d NodeHeap) Len() int           { return len(d) }
func (d NodeHeap) Less(a, b int) bool { return d[a].freq < d[b].freq }
func (d NodeHeap) Swap(a, b int)      { d[a], d[b] = d[b], d[a] }

func (d *NodeHeap) Push(x interface{}) {
	*d = append(*d, x.(Node))
}

func (d *NodeHeap) Pop() interface{} {
	old := *d
	n := len(old)
	x := old[n-1]
	*d = old[0 : n-1]
	return x
}

func Huffman(nh NodeHeap) {

}

func main() {
	nodes, err := NodesFromFile("./data.assignment3")
	if err != nil {
		log.Fatal(err)
	}
	nh := &NodeHeap{}
	for _, node := range nodes {
		*nh = append(*nh, node)
	}
	heap.Init(nh)

	for nh.Len() > 0 {
		fmt.Println(heap.Pop(nh))
	}
}

func NodesFromFile(path string) (NodeHeap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var n []Node
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		cell := scanner.Text()
		val := string(cell[1])
		freq, err := strconv.Atoi(string(cell[5]))
		if err !=  nil {
			return nil, err
		}
		n = append(n, *NewNode(val, freq))
	}
	return n, nil
}
