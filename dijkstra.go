package skiena

import (
	"container/heap"
	"fmt"
)

type searchNode struct {
	id    int
	dist  int
	index int
}

type searchNodeQueue []*searchNode

func (s searchNodeQueue) Len() int {
	return len(s)
}

func (s searchNodeQueue) Less(i, j int) bool {
	return s[i].dist < s[j].dist
}

func (s searchNodeQueue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].index = s[j].index
	s[j].index = s[i].index
}

func (s *searchNodeQueue) Push(v interface{}) {
	item := v.(*searchNode)
	item.index = len(*s)
	*s = append(*s, item)
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
	heap.Push(&pq, &searchNode{v0.id, 0, 0})
	best := make(map[int]*searchNode)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*searchNode)
		if node.id == to {
			return node.dist, nil
		}
		best[node.id] = node
		v := g.vertices[node.id]
		for _, e := range v.adjacent {
			d := node.dist + e.weight
			next, seen := best[e.to]
			if !seen {
				next = &searchNode{id: e.to, dist: d}
				heap.Push(&pq, next)
				best[next.id] = next
			} else if next.dist > d {
				next.dist = d
				heap.Fix(&pq, next.index)
			}
		}
	}
	return -1, fmt.Errorf("No path from %d to %d", from, to)
}
