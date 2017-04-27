package main

import (
	"container/heap"
	"fmt"
)

type Weights struct {
	val string
	w   float64
}

type WeightedVertex struct {
	v Vertex
	w float64
}

type Path struct {
	u string
	v string
}

type Vertex struct {
	value string
	color string
	p     string
	d     float64
}

func NewVertex(value string) *Vertex {
	return &Vertex{
		value: value,
		color: "white",
		p:     "None",
		d:     -1e2,
	}

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
	fmt.Println("Vertex:\t\tParent:\t\tColor:")
	for _, v := range g.V {
		fmt.Printf("%s\t\t%s\t\t%s\n", v.value, v.p, v.color)
	}
	fmt.Println()
}

func BFS(gr *Graph, startVert *Vertex, goal string) {
	startVert.color = "gray"
	startVert.p = "None"
	startVert.d = 0.0
	var hp VertexHeap
	hp = make([]Vertex, 0)
	heap.Init(&hp)
	heap.Push(&hp, *startVert)

	for hp.Len() > 0 {
		u := heap.Pop(&hp).(Vertex)

		for _, v := range gr.Adj[u.value] {

			tmp := gr.getV(v.v.value)
			//fmt.Println("where are you?", v)
			if tmp.color == "white" && v.w > 0 {
				tmp.color = "gray"
				tmp.d = u.d + 1
				tmp.p = u.value

				heap.Push(&hp, *tmp)
				fmt.Println("heap: ", hp)
				gr.printG()
			}
		}
		gr.setColor(u, "black")
		if u.value == goal {
			return
		}
	}
}

func getAugmentedPath(gr *Graph, goal string) []Path {
	d := make(map[string]string, 0)
	for _, v := range gr.V {
		d[v.value] = v.p
	}
	augmentedPath := make([]Path, 0)
	sink := goal
	for {
		tmp := gr.getV(sink)
		u := tmp.p
		augmentedPath = append(augmentedPath, Path{u, sink})
		sink = u
		if sink == "None" {
			return augmentedPath
		}
	}
}

func getNetFlow(Adj map[string][]Weights, augmentedPath []Path) float64 {
	var m float64
	m = 1e6
	var f float64
	for _, el := range augmentedPath {
		if el.u != "None" {
			for _, vert := range Adj[el.u] {
				if vert.val == el.v {
					f = float64(vert.w)
				}
			}
			if f < m {
				m = f
			}
		}
	}
	return m
}

func getMaxFlow(vertices []Vertex, AdjListRes map[string][]Weights, source string, goal string) float64 {
	//pAugmentedPath := make([]Path, 0)
	augmentedPath := make([]Path, 0)
	netFlow := 0.0
	maxFlow := 0.0
	for netFlow != float64(1e6) {
		G := NewGraph(vertices, AdjListRes)
		s := G.getV(source)
		//pAugmentedPath = augmentedPath[:]
		BFS(G, s, goal)
		augmentedPath = getAugmentedPath(G, goal)
		netFlow = getNetFlow(AdjListRes, augmentedPath)

		if netFlow != float64(1e6) {
			maxFlow += netFlow
		}

		for _, el := range augmentedPath {
			if el.u != "None" {
				for i := 0; i < len(AdjListRes[el.v]); i++ {
					if (AdjListRes[el.v][i].val) == el.u {
						AdjListRes[el.v][i].w += netFlow
					}
				}
				for i := 0; i < len(AdjListRes[el.u]); i++ {
					if (AdjListRes[el.u][i].val) == el.v {
						AdjListRes[el.u][i].w -= netFlow
					}
				}
			}
			fmt.Println("Augmented Path: ", augmentedPath)
			fmt.Println("New Flow = ", netFlow)
		}
	}
	return maxFlow
}

func main() {
	values := []string{"s", "v1", "v2", "v3", "v4", "t"}
	vertices := make([]Vertex, 0)
	for _, val := range values {
		vertices = append(vertices, *NewVertex(val))
	}

	/*djValues := map[string][]Weights{
		"s":  {{"v1", 16}, {"v2", 13}},
		"v1": {{"v3", 12}},
		"v2": {{"v1", 4}, {"v4", 14}},
		"v3": {{"v2", 9}, {"t", 20}},
		"v4": {{"v3", 7, "t", 4}},
		"t":  {},
	}*/

	/*
		adjValuesRes := map[string][]Weights{
			"s":  {{"v1", 16}, {"v2", 13}},
			"v1": {{"s", 0}, {"v2", 0}, {"v3", 12}},
			"v2": {{"s", 0}, {"v1", 4}, {"v3", 0}, {"v4", 14}},
			"v3": {{"v1", 0}, {"v2", 9}, {"v4", 0}, {"t", 20}},
			"v4": {{"v2", 0}, {"v3", 7}, {"t", 4}},
			"t":  {{"v3", 0}, {"v4", 0}},
		}*/

	adjValuesRes := map[string][]Weights{
		"s":  {{"v1", 16}, {"v2", 13}},
		"v1": {{"s", 0}, {"v2", 10}, {"v3", 12}},
		"v2": {{"s", 0}, {"v1", 4}, {"v3", 0}, {"v4", 14}},
		"v3": {{"v1", 0}, {"v2", 9}, {"v4", 0}, {"t", 20}},
		"v4": {{"v2", 0}, {"v3", 7}, {"t", 4}},
		"t":  {{"v3", 0}, {"v4", 0}},
	}

	source := "s"
	goal := "t"
	maxFlow := getMaxFlow(vertices, adjValuesRes, source, goal)
	fmt.Println("Max flow: ", maxFlow)

}

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
