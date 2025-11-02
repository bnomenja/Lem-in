package functions

import (
	"container/heap"
	"fmt"
	"math"
)

// Item pour la priority queue
type Item struct {
	Name     string
	Priority int
	Index    int
}

// PriorityQueue implémente heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func HasDuplicateRoomAcrossPaths(paths [][]string) bool {
	seen := make(map[string]int) // map salle -> index du premier path où elle apparaît

	for i, path := range paths {
		for _, room := range path {
			if prev, exists := seen[room]; exists && prev != i {
				// La même salle est dans un autre path
				return true
			}
			seen[room] = i
		}
	}
	return false
}

func Suurballe(farm *Farm) []Path {
	start := farm.SpecialRooms["start"]
	end := farm.SpecialRooms["end"]
	paths := []Path{}
	n := 0

	for {

		dist, path := Dijkstra(farm, start, end)
		n++

		fmt.Println(n)
		if len(path) == 0 || path[0] != start || n == 32 {
			break
		}

		UpdateGraph(farm, path, dist)

		paths = append(paths, path)

	}

	paths = MergePaths(paths, start, end)

	if HasDuplicateRoomAcrossPaths(paths) {
	}
}

func UpdateGraph(farm *Farm, shortest Path, dist map[string]int) {
	newEdges := make(map[string]Edge)

	for key, edge := range farm.Edges {
		from, to := edge.From, edge.To
		edge.Weight = edge.Weight + dist[from] - dist[to]
		newEdges[key] = edge
	}

	for i := 0; i < len(shortest)-1; i++ {
		from, to := shortest[i], shortest[i+1]
		delete(newEdges, from+"-"+to)
		newEdges[to+"-"+from] = Edge{From: to, To: from, Weight: 0}
	}

	farm.Edges = newEdges
}

func buildPathFromParents(parent map[string]string, start, end string) Path {
	path := Path{}
	for room := end; room != ""; room = parent[room] {
		path = append([]string{room}, path...)
		if room == start {
			break
		}
	}
	return path
}

func Dijkstra(farm *Farm, start, end string) (map[string]int, Path) {
	dist := make(map[string]int)
	parent := make(map[string]string)
	visited := make(map[string]bool)

	for name := range farm.Rooms {
		dist[name] = math.MaxInt32
	}

	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Name: start, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.Name

		if visited[node] {
			continue
		}
		visited[node] = true

		if node == end {
			break
		}

		for neighborName := range farm.Rooms[node].Links {
			key := node + "-" + neighborName
			edge, exists := farm.Edges[key]
			if !exists {
				continue
			}

			newDist := dist[node] + edge.Weight
			if newDist < dist[neighborName] {
				dist[neighborName] = newDist
				parent[neighborName] = node
				heap.Push(pq, &Item{Name: neighborName, Priority: newDist})
			}
		}
	}

	return dist, buildPathFromParents(parent, start, end)
}

// MergePaths transforme la liste de chemins trouvés en un ensemble de chemins
// node-disjoints (sauf start/end). Elle utilise une réduction node->(in,out)
// et Edmonds-Karp pour trouver un flot maximal, puis décompose le flot en chemins.
func MergePaths(paths []Path, start, end string) []Path {
	// 1) construire l'ensemble des arêtes (union des paths)
	edges := map[string]map[string]bool{}
	nodes := map[string]bool{}
	for _, p := range paths {
		for i := 0; i < len(p)-1; i++ {
			u, v := p[i], p[i+1]
			if edges[u] == nil {
				edges[u] = map[string]bool{}
			}
			edges[u][v] = true
			nodes[u] = true
			nodes[v] = true
		}
	}

	// helper pour split nodes
	inName := func(n string) string { return n + "#in" }
	outName := func(n string) string { return n + "#out" }

	// 2) construire capacités (residual capacities)
	cap := map[string]map[string]int{} // cap[u][v] = capacity
	addCap := func(u, v string, c int) {
		if cap[u] == nil {
			cap[u] = map[string]int{}
		}
		cap[u][v] = cap[u][v] + c
	}

	// 2a) node splitting: for each node x != start,end add x_in->x_out cap=1
	// start/end get large capacity (len(paths) or large)
	maxCap := len(paths) + 5
	if maxCap < 10 {
		maxCap = 10
	}
	for n := range nodes {
		if n == start || n == end {
			addCap(inName(n), outName(n), maxCap)
		} else {
			addCap(inName(n), outName(n), 1)
		}
	}

	// 2b) for each union edge u->v add edge u_out -> v_in cap=1
	for u, m := range edges {
		for v := range m {
			addCap(outName(u), inName(v), 1)
		}
	}

	// 3) Edmonds-Karp (BFS augmenting paths) on residual network
	// residual initially = cap
	res := map[string]map[string]int{}
	for u, m := range cap {
		if res[u] == nil {
			res[u] = map[string]int{}
		}
		for v, c := range m {
			res[u][v] = c
		}
	}

	src := outName(start) // source is start_out
	sink := inName(end)   // sink is end_in

	// BFS to find augmenting path, returns parent map
	bfs := func() map[string]string {
		parent := map[string]string{}
		q := []string{src}
		visited := map[string]bool{src: true}
		for len(q) > 0 {
			x := q[0]
			q = q[1:]
			if x == sink {
				break
			}
			for y, c := range res[x] {
				if c > 0 && !visited[y] {
					visited[y] = true
					parent[y] = x
					q = append(q, y)
				}
			}
		}
		if !visited[sink] {
			return nil
		}
		return parent
	}

	// augment flows until no augmenting path
	for {
		parent := bfs()
		if parent == nil {
			break
		}
		// find bottleneck
		bottleneck := int(^uint(0) >> 1) // max int
		for v := sink; v != src; v = parent[v] {
			u := parent[v]
			if res[u][v] < bottleneck {
				bottleneck = res[u][v]
			}
		}
		// apply
		for v := sink; v != src; v = parent[v] {
			u := parent[v]
			// decrease forward, increase backward
			res[u][v] -= bottleneck
			if res[v] == nil {
				res[v] = map[string]int{}
			}
			res[v][u] += bottleneck
		}
	}

	// 4) décomposer le flot en chemins unitaires :
	// on parcourt le graphe des arêtes qui ont été saturées dans le sens forward
	// (i.e. capacité utilisée = original cap - residual forward > 0)
	// pour faciliter l'extraction, calculons flow[u][v] = originalCap - res[u][v]
	flow := map[string]map[string]int{}
	for u, m := range cap {
		for v, c := range m {
			used := c - res[u][v] // si res[u][v] absent => 0, handled by zero value
			if used > 0 {
				if flow[u] == nil {
					flow[u] = map[string]int{}
				}
				flow[u][v] = used
			}
		}
	}

	// extraction : tant qu'il existe un chemin start_out -> end_in dans flow, on le retire
	var final []Path
	for {
		// find a path in flow using DFS/BFS
		parent := map[string]string{}
		q := []string{src}
		visited := map[string]bool{src: true}
		found := false
		for len(q) > 0 && !found {
			x := q[0]
			q = q[1:]
			if x == sink {
				found = true
				break
			}
			for y, f := range flow[x] {
				if f > 0 && !visited[y] {
					visited[y] = true
					parent[y] = x
					q = append(q, y)
					if y == sink {
						found = true
						break
					}
				}
			}
		}
		if !found {
			break
		}
		// reconstruire le chemin de src->sink
		// et décrémenter flow le long du chemin (consommer 1 unité)
		// recueillir les noms de salles originaux (convertir in/out)
		stack := []string{}
		for v := sink; v != src; v = parent[v] {
			stack = append(stack, v)
		}
		stack = append(stack, src)
		// stack contains nodes from sink...src, reverse it
		for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
			stack[i], stack[j] = stack[j], stack[i]
		}

		// consume flows and build path of original node names:
		pathNodes := []string{}
		// stack is [start_out, ..., end_in]
		for i := 0; i < len(stack)-1; i++ {
			u := stack[i]
			v := stack[i+1]
			// decrement flow
			flow[u][v]--
			// derive original node when crossing out->in or in->out edges
			// we only append original node names when we enter an "in" node or start
			// rule: when node is start_out => append start; when node is something_in => append nodeName
			if i == 0 {
				// first node is start_out -> append start
				pathNodes = append(pathNodes, start)
			}
			// if v ends with "#in", append original node v without suffix
			if len(v) > 3 && v[len(v)-3:] == "#in" {
				orig := v[:len(v)-3]
				// except if orig == start (already added) we add orig only when it's not start
				if orig != start {
					pathNodes = append(pathNodes, orig)
				}
			}
			// if v == sink (end_in), we will ensure end appended below
			if v == sink {
				// append end if not already appended
				if len(pathNodes) == 0 || pathNodes[len(pathNodes)-1] != end {
					pathNodes = append(pathNodes, end)
				}
			}
		}

		final = append(final, Path(pathNodes))
	}

	return final
}
