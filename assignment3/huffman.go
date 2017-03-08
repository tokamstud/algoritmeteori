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
	left *Node
	right *Node
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
var codeTable map[string]string

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

func goFuckYourself() string {
	return "Go fuck yourself"
}

func Huffman(nh *NodeHeap) Node {
	n := nh.Len()
	qnh := nh
	for i := 1; i < n; i++ {
		z := &Node{}
		x := heap.Pop(qnh).(Node)
		y := heap.Pop(qnh).(Node)
		z.left = &x
		z.right = &y
		z.freq = x.freq + y.freq
		heap.Push(qnh, *z)
	}
	return heap.Pop(qnh).(Node)
}

func main() {
	nodes, err := NodesFromFile("./data.assignment3")
	if err != nil {
		log.Fatal(err)
	}
	raw := make(map[string]int)

	nh := &NodeHeap{}
	var bits int

	for _, node := range nodes {
		*nh = append(*nh, node)
		raw[node.val] = node.freq
		bits+=node.freq*8
	}
	heap.Init(nh)

	// createing huffman tree
	qnh := Huffman(nh)

	codeTable = make(map[string]string)
	getCodes(&qnh, "", codeTable)

	fmt.Println("Before compression", bits)
	bits = 0
	for key, val := range codeTable {
		fmt.Println(val,raw[val], key)
		bits+=len(key)*raw[val]
	}
	fmt.Println("After compression", bits)
}

func getCodes(root *Node, code string, codeTable map[string]string) map[string]string {
	if root != nil {
		if root.left != nil {
			getCodes(root.left, code+"0", codeTable)
		}
		if root.right != nil {
			getCodes(root.right, code+"1", codeTable)
		}
		if root.right == nil && root.right == nil {
			codeTable[code] = root.val
		}
	}
	return codeTable

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
