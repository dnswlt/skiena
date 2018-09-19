package skiena

import "fmt"

// Visitor is an interface grouping functions called during a dfs graph traversal
type Visitor interface {
	// Stopped should yield true to interrupt the current graph traversal. It is called before PreprocessVertex.
	Stopped() bool
	// PreprocessVertex is called when a vertex is encountered for the first time, before its outgoing edges are processed
	PreprocessVertex(id int)
	// ProcessEdge is called for each outgoing edge of the currently processed vertex
	ProcessEdge(e *edge)
	// PostprocessVertex is called
	PostprocessVertex(id int)
}

// DepthFirstSearch performs a depth-first search on graph g, starting at node start.
func DepthFirstSearch(start int, g *Graph, viz Visitor) {
	m := make(map[int]bool)
	dfs(start, g, viz, m)
}

func dfs(start int, g *Graph, viz Visitor, seen map[int]bool) {
	v, ok := g.vertices[start]
	if !ok {
		fmt.Printf("No such vertex %d\n", start)
		return
	}
	if viz.Stopped() {
		return
	}
	viz.PreprocessVertex(v.id)
	seen[v.id] = true
	for _, e := range v.adjacent {
		viz.ProcessEdge(e)
		if seen[e.to] {
			continue
		}
		dfs(e.to, g, viz, seen)
	}
	viz.PostprocessVertex(v.id)
}
