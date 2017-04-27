package main

import (
	"container/heap"
	"fmt"
)

type Vertex struct {
	value string
	color string
	d     float64
	p     string
	f     float64
}

func NewVertex(value string) *Vertex {
	return &Vertex{
		value: value,
		color: "white",
		d:     -1e2,
		p:     "None",
		f:     -1.0,
	}

}

type Graph struct {
	V   []Vertex
	Adj map[string][]Vertex
}

func NewGraph(vertices []Vertex, adjacencyList map[string][]string) *Graph {
	g := &Graph{}
	g.V = make([]Vertex, 0)
	g.Adj = make(map[string][]Vertex, 0)
	for _, v := range vertices {
		g.V = append(g.V, v)
	}
	for key := range adjacencyList {
		objects := make([]Vertex, 0)
		values := adjacencyList[key]
		for _, v := range values {
			objects = append(objects, *g.getV(v))
		}
		g.Adj[key] = objects
	}
	return g
}

func (g *Graph) getV(value string) *Vertex {
	for i := 0; i < len(g.V); i++ {
		if g.V[i].value == value {
			return &g.V[i]
		}
	}
	return &Vertex{}
}

func (g *Graph) setColor(val Vertex, col string) {
	for i := 0; i < len(g.V); i++ {
		if g.V[i].value == val.value {
			g.V[i].color = col
		}
	}
}

func (g *Graph) printG() {
	fmt.Println("Vertex:\t\tParent:\t\tDistance:\tTime:\tColor:")
	for _, v := range g.V {
		fmt.Printf("%s\t\t%s\t\t%.0f\t\t%.0f\t%s\n", v.value, v.p, v.d, v.f, v.color)
	}
	fmt.Println()
}

func DFS(gr *Graph) {
	time := 0.0
	for i := 0; i < len(gr.V); i++ {
		if gr.V[i].color == "white" {
			time = dfsVisit(gr, &gr.V[i], time)
		}
	}
}

func dfsVisit(gr *Graph, vert *Vertex, time float64) float64 {
	gr.printG()
	time++
	vert.d = time
	vert.color = "gray"
	for _, v := range gr.Adj[vert.value] {
		tmp := gr.getV(v.value)
		if tmp.color == "white" {
			tmp.p = vert.value
			time = dfsVisit(gr, tmp, time)
		}
	}
	vert.color = "black"
	time++
	vert.f = time
	return time
}

func BFS(gr *Graph, startVert *Vertex) {
	startVert.color = "gray"
	startVert.d = 0.0
	startVert.p = "None"
	var hp VertexHeap
	hp = make([]Vertex, 0)
	heap.Init(&hp)
	heap.Push(&hp, *startVert)

	for hp.Len() > 0 {
		u := heap.Pop(&hp).(Vertex)

		for _, v := range gr.Adj[u.value] {

			tmp := gr.getV(v.value)
			fmt.Println("where are you?", v)
			if tmp.color == "white" {
				tmp.color = "gray"
				tmp.d = u.d + 1
				tmp.p = u.value

				heap.Push(&hp, *tmp)
				fmt.Println("heap: ", hp)
				gr.printG()
			}
		}
		gr.setColor(u, "black")
	}
}

func main() {

	// Undirected graph for BFS
	values := []string{"r", "s", "t", "u", "v", "w", "x", "y"}
	vertices := make([]Vertex, 0)
	for _, val := range values {
		vertices = append(vertices, *NewVertex(val))
	}

	adjValues := map[string][]string{
		"r": {"s", "v"},
		"s": {"r", "w"},
		"t": {"u", "w", "x"},
		"u": {"t", "x", "y"},
		"v": {"r"},
		"w": {"s", "t", "x"},
		"x": {"t", "u", "w", "y"},
		"y": {"u", "x"},
	}

	G := NewGraph(vertices, adjValues)

	G.printG()

	s := G.getV("s")

	BFS(G, s)

	G.printG()

	// Directed graph for DFS
	values2 := []string{"u", "v", "w", "x", "y", "z"}
	vertices2 := make([]Vertex, 0)
	for _, val := range values2 {
		vertices2 = append(vertices2, *NewVertex(val))
	}

	adjValues2 := map[string][]string{
		"u": {"x", "v"},
		"v": {"y"},
		"w": {"y", "z"},
		"x": {"v"},
		"y": {"x"},
		"z": {"z"},
	}

	G2 := NewGraph(vertices2, adjValues2)

	G2.printG()

	DFS(G2)

	G2.printG()
}

// Methods to form Heap

type VertexHeap []Vertex

func (vh VertexHeap) Len() int { return len(vh) }

func (vh VertexHeap) Less(i, j int) bool {

	return vh[i].d < vh[j].d

}
func (vh VertexHeap) Swap(i, j int) { vh[i], vh[j] = vh[j], vh[i] }

func (vh *VertexHeap) Push(x interface{}) {
	*vh = append(*vh, x.(Vertex))
}

func (vh *VertexHeap) Pop() interface{} {
	old := *vh
	n := len(old)
	x := old[n-1]
	*vh = old[0 : n-1]
	return x
}
