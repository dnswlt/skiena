package skiena

import "testing"

func TestSize(t *testing.T) {
	g := MakeGraph([]int{1, 2, 1, 3, 2, 4, 2, 5, 3, 6, 3, 7})
	if g.Size() != 7 {
		t.Errorf("Incredible graph size %d\n", g.Size())
	}
}

func TestReadDirectedGraph(t *testing.T) {
	g, err := ReadDirectedGraph("samplegraph.json")
	if err != nil {
		t.Error("Couldn't read sample graph: ", err)
	}
	if g.Size() != 3 {
		t.Errorf("Unexpected graph size: %d", g.Size())
	}
}
