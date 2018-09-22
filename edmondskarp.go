package skiena

type residualEdge struct {
	from     int
	to       int
	flow     int
	capacity int
	rev      *residualEdge
	edge     *edge
}

// Create a residual graph of g, i.e. one in which there's an additional reverse edge for each edge of g.
func makeResidual(g *Graph) map[int][]*residualEdge {
	m := make(map[int][]*residualEdge)
	for _, v := range g.vertices {
		for _, e := range v.adjacent {
			fwdEdge := &residualEdge{v.id, e.to, 0, e.weight, nil, e}
			revEdge := &residualEdge{e.to, v.id, 0, 0, fwdEdge, nil}
			fwdEdge.rev = revEdge
			m[v.id] = append(m[v.id], fwdEdge)
			m[e.to] = append(m[e.to], revEdge)
		}
	}
	return m
}

// MaximumFlow returns the maximum flow possible in graph g from node from to node to.
func MaximumFlow(g *Graph, from, to int) (int, map[*edge]int) {
	src, ok := g.vertices[from]
	if !ok {
		return 0, nil
	}
	dest, ok := g.vertices[to]
	if !ok {
		return 0, nil
	}
	// build residual graph (whatever residual means...)
	r := makeResidual(g)
	flow := 0
	for i := 0; i < 10; i++ {
		// find augmenting path
		q := []int{src.id}
		pred := make(map[int]*residualEdge)
	bfs:
		for len(q) > 0 {
			i := q[0]
			q = q[1:]
			for _, re := range r[i] {
				_, seen := pred[re.to]
				if !seen && re.to != src.id && re.capacity > re.flow {
					pred[re.to] = re
					if re.to == to {
						break bfs
					}
					q = append(q, re.to)
				}
			}
		}
		e, found := pred[to]
		if !found {
			// no augmenting path found
			break
		}
		// find "delta" by which to increment flow
		d := 0
		for e != nil {
			if d == 0 || e.capacity-e.flow < d {
				d = e.capacity - e.flow
			}
			e = pred[e.from]
		}
		// increase flow
		e = pred[dest.id]
		for e != nil {
			e.flow += d
			e.rev.flow -= d
			e = pred[e.from]
		}
		flow += d
	}
	ef := make(map[*edge]int)
	for _, res := range r {
		for _, re := range res {
			if re.edge != nil {
				ef[re.edge] = re.flow
			}
		}
	}
	return flow, ef
}
