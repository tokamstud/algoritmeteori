package main

import (
	"container/heap"
	"fmt"
)

type Weights struct {
	val string
	w   int
}

type WeightedVertex struct {
	v Vertex
	w int
}

type Vertex struct {
	value string
	d     float64
	p     string
}

func NewVertex(value string) *Vertex {
	return &Vertex{
		value: value,
		d:     1e3,
		p:     "None",
	}
}

func (v *Vertex) printV() {
	fmt.Println(v.value, v.d, v.p)
}

type Graph struct {
	V   []Vertex
	Adj map[string][]WeightedVertex
}

func NewGraph(vertices []Vertex, adjacencyList map[string][]Weights) *Graph {
	g := &Graph{}
	g.V = make([]Vertex, 0)
	g.Adj = make(map[string][]WeightedVertex, 0)
	for _, v := range vertices {
		g.V = append(g.V, v)
	}
	for key := range adjacencyList {
		objects := make([]WeightedVertex, 0)
		values := adjacencyList[key]
		for _, v := range values {
			tmp := *g.getV(v.val)
			newVert := WeightedVertex{tmp, v.w}
			objects = append(objects, newVert)
		}
		g.Adj[key] = objects
	}
	return g
}

func (g *Graph) reset() {
	for i := 0; i < len(g.V); i++ {
		g.V[i].d = 1e3
		g.V[i].p = "None"
	}
}

func (g *Graph) getV(value string) *Vertex {
	for i := 0; i < len(g.V); i++ {
		if g.V[i].value == value {
			return &g.V[i]
		}
	}
	return &Vertex{}
}

func (g *Graph) getAllVertices() []Vertex {
	// Maybe use pointer if there will be problems!!!
	vrs := make([]Vertex, 0)
	for i := 0; i < len(g.V); i++ {
		vrs = append(vrs, g.V[i])
	}
	return vrs
}

func (g *Graph) printG() {
	fmt.Println("Vertex:\t\tParent:\t\tDistance:")
	for _, v := range g.V {
		fmt.Printf("%s\t\t%s\t\t%.0f\n", v.value, v.p, v.d)
	}
	fmt.Println()
}

func dijkstra(gr *Graph, startVert *Vertex) {

	startVert.d = 0.0
	S := make([]Vertex, 0)

	var Q WVertexHeap
	Q = make([]Vertex, 0)
	heap.Init(&Q)

	for _, v := range gr.V {
		tmp := gr.getV(v.value)
		heap.Push(&Q, *tmp)
	}
	//fmt.Println("Initial heap: ", Q)
	for Q.Len() > 0 {
		u := heap.Pop(&Q).(Vertex)

		S = append(S, u)
		gr.printG()

		for _, el := range gr.Adj[u.value] {
			alt := u.d + float64(el.w)
			temp := gr.getV(el.v.value)
			if alt < temp.d {
				temp.d = alt
				temp.p = u.value
				heap.Push(&Q, *temp)
				//fmt.Println("Heap: ", Q)
			}
		}
	}
	fmt.Println("S: ", S)
}

func main() {
	values := []string{"A", "B", "C", "D", "E"}
	vertices := make([]Vertex, 0)
	for _, val := range values {
		vertices = append(vertices, *NewVertex(val))
	}

	adjValues := map[string][]Weights{
		"A": {{"B", 10}, {"C", 3}},
		"B": {{"C", 1}, {"D", 2}},
		"C": {{"D", 8}, {"E", 2}},
		"D": {{"E", 7}},
		"E": {{"D", 9}},
	}

	G := NewGraph(vertices, adjValues)

	G.printG()

	s := G.getV("A")
	dijkstra(G, s)
	G.printG()
}

// heap setup
type WVertexHeap []Vertex

func (wvh WVertexHeap) Len() int { return len(wvh) }

func (wvh WVertexHeap) Less(i, j int) bool {
	return wvh[i].d < wvh[j].d
}
func (wvh WVertexHeap) Swap(i, j int) { wvh[i], wvh[j] = wvh[j], wvh[i] }

func (wvh *WVertexHeap) Push(x interface{}) {
	*wvh = append(*wvh, x.(Vertex))
}

func (wvh *WVertexHeap) Pop() interface{} {
	old := *wvh
	n := len(old)
	x := old[n-1]
	*wvh = old[0 : n-1]
	return x
}
