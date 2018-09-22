package skiena

import "strings"

// MatchCourses tries to find a minimum number of courses a student has to attend
// to obtain a degree. Each requirement has the form "\d[A-Z]+", e.g. "2ABCD", indicating
// that the student has to attend 2 courses out of A, B, C, D to satisfy the requirement.
// The function returns the sequence of courses the student should attend.
func MatchCourses(requirements []string) string {
	g := NewGraph()
	const (
		source = -1
		sink   = -2
	)
	g.AddVertex(source)
	g.AddVertex(sink)
	rID := sink - 1
	numCourses := 0
	for _, req := range requirements {
		n := int(req[0] - '0')
		if len(req[1:]) < n {
			return ""
		}
		numCourses += n
		// Add requirement vertex
		vr := g.AddVertex(rID)
		rID--
		g.AddDirected(source, vr.id, n)
		for i := 1; i < len(req); i++ {
			vID := int(req[i])
			if _, found := g.vertices[vID]; !found {
				g.AddVertex(vID)
				// Add edge course->sink
				g.AddDirected(vID, sink, 1)
			}
			// Add edge req->course
			g.AddDirected(vr.id, vID, 1)
		}
	}
	flow, edgeFlows := MaximumFlow(g, source, sink)
	if flow != numCourses {
		return ""
	}
	var buf strings.Builder
	for _, re := range g.vertices[source].adjacent {
		for _, ce := range g.vertices[re.to].adjacent {
			if edgeFlows[ce] > 0 {
				buf.WriteByte(byte(ce.to))
			}
		}
	}
	return buf.String()
}
