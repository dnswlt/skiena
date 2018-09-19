package skiena

import (
	"fmt"
	"testing"
)

type countingVisitor struct {
	countPreprocess  int
	preprocessed     []int
	countEdge        int
	edges            []*edge
	countPostprocess int
	postprocessed    []int
}

func (v *countingVisitor) PreprocessVertex(id int) {
	v.countPreprocess++
	v.preprocessed = append(v.preprocessed, id)
}

func (v *countingVisitor) PostprocessVertex(id int) {
	v.countPostprocess++
	v.postprocessed = append(v.postprocessed, id)
}

func (v *countingVisitor) ProcessEdge(e *edge) {
	v.countEdge++
	v.edges = append(v.edges, e)
}

func (v *countingVisitor) Stopped() bool {
	return false
}

func TestDfs(t *testing.T) {
	g := MakeGraph([]int{1, 2, 1, 3, 2, 4, 2, 5, 3, 6, 3, 7})
	viz := countingVisitor{}
	DepthFirstSearch(1, g, &viz)
	if viz.countEdge != 6 {
		t.Errorf("Funny edge count: %d\n", viz.countEdge)
	}
	if viz.countPreprocess != 7 {
		t.Errorf("Funny preprocess count: %d\n", viz.countPreprocess)
	}
	if viz.countPostprocess != 7 {
		t.Errorf("Funny postprocess count: %d\n", viz.countPostprocess)
	}
	if fmt.Sprint(viz.preprocessed) != "[1 2 4 5 3 6 7]" {
		t.Error("Expected different postprocessed", viz.preprocessed)
	}
	if fmt.Sprint(viz.postprocessed) != "[4 5 2 6 7 3 1]" {
		t.Error("Expected different postprocessed", viz.postprocessed)
	}
}

func isTopologicallySorted(edges []*edge, vs []int) bool {
	pos := make(map[int]int)
	for i, v := range vs {
		pos[v] = i
	}
	for _, e := range edges {
		p1, ok := pos[e.from]
		if !ok {
			return false
		}
		p2, ok := pos[e.to]
		if !ok {
			return false
		}
		if p1 <= p2 {
			return false
		}
	}
	return true
}

func TestDfsDAG(t *testing.T) {
	g := MakeGraph([]int{1, 2, 1, 3, 2, 4, 2, 5, 3, 6, 3, 7, 4, 8, 5, 8, 5, 9, 6, 9, 6, 10, 7, 10, 8, 11, 9, 11, 9, 12, 10, 12, 11, 13, 12, 13})
	viz := countingVisitor{}
	DepthFirstSearch(1, g, &viz)
	if viz.countEdge != 18 {
		t.Errorf("Funny edge count: %d\n", viz.countEdge)
	}
	if viz.countPreprocess != 13 {
		t.Errorf("Funny preprocess count: %d\n", viz.countPreprocess)
	}
	if viz.countPostprocess != 13 {
		t.Errorf("Funny postprocess count: %d\n", viz.countPostprocess)
	}
	if !isTopologicallySorted(viz.edges, viz.postprocessed) {
		t.Error("Not topologically sorted", viz.postprocessed)
	}
}
