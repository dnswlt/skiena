package skiena

import "testing"

func TestEdmondsKarp(t *testing.T) {
	g := MakeGraph([]int{1, 2, 1, 3, 2, 4, 2, 5, 3, 5, 4, 6, 5, 6})
	f := MaximumFlow(g, 1, 6)
	if f != 2 {
		t.Errorf("Wrong flow: %d", f)
	}
}

func TestEdmondsKarpWeighted(t *testing.T) {
	g := MakeWeightedGraph([]int{
		1, 2, 2,
		1, 3, 1,
		2, 4, 1,
		2, 5, 1,
		3, 5, 1,
		4, 6, 1,
		5, 6, 2,
	})
	f := MaximumFlow(g, 1, 6)
	if f != 3 {
		t.Errorf("Wrong flow: %d", f)
	}
}
