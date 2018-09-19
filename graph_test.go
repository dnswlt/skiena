package skiena

import "testing"

func TestSize(t *testing.T) {
	g := MakeGraph([]int{1, 2, 1, 3, 2, 4, 2, 5, 3, 6, 3, 7})
	if g.Size() != 7 {
		t.Errorf("Incredible graph size %d\n", g.Size())
	}
}
