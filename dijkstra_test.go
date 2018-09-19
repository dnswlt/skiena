package skiena

import "testing"

func TestDijkstra(t *testing.T) {
	g := NewGraph()
	for i := 1; i <= 1000; i++ {
		g.AddVertex(i)
	}
	for i := 1; i <= 999; i++ {
		g.AddDirected(i, 1000, 2000-2)
		g.AddDirected(i, i+1, 1)
	}
	dist, err := ShortestPathDijkstra(g, 1, 1000)
	if err != nil || dist != 999 {
		t.Errorf("Baaad dist: %d", dist)
	}
}

func BenchmarkDijkstra(b *testing.B) {
	g := NewGraph()
	size := 200000
	for i := 1; i <= size; i++ {
		g.AddVertex(i)
	}
	for i := 1; i <= size-1; i++ {
		g.AddDirected(i, size, size*2)
		g.AddDirected(i, i+1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShortestPathDijkstra(g, 1, size)
	}
}
