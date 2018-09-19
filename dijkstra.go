package skiena

import (
	"container/heap"
	"fmt"
)

type searchNode struct {
	id   int
	dist int
}

type searchNodeQueue []searchNode

func (s searchNodeQueue) Len() int {
	return len(s)
}

func (s searchNodeQueue) Less(i, j int) bool {
	return s[i].dist < s[j].dist
}

func (s searchNodeQueue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *searchNodeQueue) Push(v interface{}) {
	*s = append(*s, v.(searchNode))
}

func (s *searchNodeQueue) Pop() interface{} {
	l := len(*s)
	old := *s
	item := old[l-1]
	*s = old[:l-1]
	return item
}

// ShortestPathDijkstra returns the length of the shortest path from from to to.
func ShortestPathDijkstra(g *Graph, from int, to int) (int, error) {
	v0 := g.vertices[from]
	pq := searchNodeQueue{}
	heap.Push(&pq, searchNode{v0.id, 0})
	best := make(map[int]int)
	maxSize := 0
	for pq.Len() > 0 {
		if pq.Len() > maxSize {
			maxSize = pq.Len()
		}
		node := heap.Pop(&pq).(searchNode)
		v := g.vertices[node.id]
		if v.id == to {
			// fmt.Printf("maxSize: %d\n", maxSize)
			return node.dist, nil
		}
		best[v.id] = node.dist
		for _, e := range v.adjacent {
			if d, seen := best[e.to]; !seen || d > node.dist+e.weight {
				heap.Push(&pq, searchNode{e.to, node.dist + e.weight})
			}
		}
	}
	return -1, fmt.Errorf("No path from %d to %d", from, to)
}
