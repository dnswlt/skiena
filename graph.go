package skiena

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type edge struct {
	from   int
	to     int
	weight int
}

type vertex struct {
	id       int
	adjacent []*edge
}

// Graph is a directed graph made of vertices that have adjacency lists
type Graph struct {
	vertices map[int]*vertex
}

// NewGraph creates a new, empty graph
func NewGraph() *Graph {
	return &Graph{make(map[int]*vertex)}
}

// Size yields the number of vertices in the graph
func (g *Graph) Size() int {
	return len(g.vertices)
}

// AddVertex adds a vertex to the graph. If the vertex already exists, this function does nothing.
func (g *Graph) AddVertex(id int) {
	_, found := g.vertices[id]
	if found {
		return
	}
	g.vertices[id] = &vertex{id: id}
}

// AddUndirected adds an undirected edge src--tgt to the graph
func (g *Graph) AddUndirected(src int, tgt int, weight int) {
	s, ok := g.vertices[src]
	if !ok {
		return
	}
	t, ok := g.vertices[tgt]
	if !ok {
		return
	}
	s.adjacent = append(s.adjacent, &edge{src, tgt, weight})
	t.adjacent = append(t.adjacent, &edge{tgt, src, weight})
}

// AddDirected adds a directed edge src->tgt to the graph
func (g *Graph) AddDirected(src int, tgt int, weight int) {
	s, ok := g.vertices[src]
	if !ok {
		return
	}
	_, ok = g.vertices[tgt]
	if !ok {
		return
	}
	s.adjacent = append(s.adjacent, &edge{src, tgt, weight})
}

// MakeGraph creates a simple graph with edge weight 1 from the slice of pairs of vertex ids.
func MakeGraph(vertexPairs []int) *Graph {
	g := NewGraph()
	for i := 0; i < len(vertexPairs)-1; i += 2 {
		g.AddVertex(vertexPairs[i])
		g.AddVertex(vertexPairs[i+1])
		g.AddDirected(vertexPairs[i], vertexPairs[i+1], 1)
	}
	return g
}

// MakeWeightedGraph creates a weighted graph from triples of ints (from, to, weight), all in one slice
func MakeWeightedGraph(vals []int) *Graph {
	g := NewGraph()
	for i := 0; i < len(vals)-2; i += 3 {
		g.AddVertex(vals[i])
		g.AddVertex(vals[i+1])
		g.AddDirected(vals[i], vals[i+1], vals[i+2])
	}
	return g
}

// ReadDirectedGraph reads a graph from a JSON file, format {"edges": [[1,2,1],[1,3,1]]}
func ReadDirectedGraph(filename string) (*Graph, error) {
	var obj struct {
		Edges [][]int `json:"edges"`
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	g := NewGraph()
	for i := 0; i < len(obj.Edges); i++ {
		l := len(obj.Edges[i])
		if l == 2 || l == 3 {
			g.AddVertex(obj.Edges[i][0])
			g.AddVertex(obj.Edges[i][1])
			w := 1
			if l == 3 {
				w = obj.Edges[i][2]
			}
			g.AddDirected(obj.Edges[i][0], obj.Edges[i][1], w)
		} else {
			return nil, fmt.Errorf("Invalid edge value %v", obj.Edges[i])
		}
	}
	return g, nil
}
