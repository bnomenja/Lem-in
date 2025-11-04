package functions

import (
	"math"
	"sort"
)

type Node struct {
	Name     string
	Priority int
}

type queue []Node

func Suurballe(farm *Farm) []Path {
	start := farm.SpecialRooms["start"]
	end := farm.SpecialRooms["end"]
	paths := []Path{}

	for {

		dist, path := FindPaths(farm, start, end)
		if path == nil {
			break
		}

		UpdateGraph(farm, path, dist)

		paths = append(paths, path)
	}

	if len(paths) == 0 {
		return paths
	}

	paths = MergePaths(paths, start, end)

	return paths
}

func MergePaths(paths []Path, start, end string) []Path {
	return paths
}

func UpdateGraph(farm *Farm, shortest Path, dist map[string]int) {
	newEdges := make(map[string]Edge)
	seen := map[string]bool{}

	for key, edge := range farm.Edges {
		if seen[key] {
			continue
		}

		from, to := edge.From, edge.To
		Weight1 := edge.Weight + dist[from] - dist[to]
		reverse := to + "-" + from

		if _, exist := farm.Edges[reverse]; exist {
			weight2 := edge.Weight + dist[to] - dist[from]

			if Weight1 > weight2 {
				Weight1 = weight2
			}

			if Weight1 < 0 {
				Weight1 = 0
			}

			seen[reverse] = true
			newEdges[reverse] = Edge{From: to, To: from, Weight: Weight1}

		}

		if Weight1 < 0 {
			Weight1 = 0
		}

		edge.Weight = Weight1
		newEdges[key] = edge
	}

	for i := 0; i < len(shortest)-1; i++ {
		from, to := shortest[i], shortest[i+1]
		delete(newEdges, from+"-"+to)
	}

	farm.Edges = newEdges
}

func (queue *queue) Add(room Node) {
	*queue = append(*queue, room)

	sort.Slice(*queue, func(i, j int) bool {
		return (*queue)[i].Priority < (*queue)[j].Priority
	})
}

func (queue *queue) Pop() Node {
	room := (*queue)[0]
	*queue = (*queue)[1:]

	return room
}

func Dijkstra(farm *Farm, start, end string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	parent := make(map[string]string)
	visited := make(map[string]bool)

	for name := range farm.Rooms {
		dist[name] = math.MaxInt
	}
	dist[start] = 0

	queue := queue{}
	queue.Add(Node{Name: start, Priority: 0})

	for len(queue) > 0 {
		node := queue.Pop()
		current, Value := node.Name, node.Priority

		visited[current] = true

		if dist[current] < Value {
			continue
		}

		for _, neighbor := range farm.Rooms[current].Links {
			if visited[neighbor.Name] {
				continue
			}

			key := current + "-" + neighbor.Name
			edge, exists := farm.Edges[key]
			if !exists {
				continue
			}

			newdist := dist[current] + edge.Weight

			if newdist < dist[neighbor.Name] {
				parent[neighbor.Name] = current
				dist[neighbor.Name] = newdist
				queue.Add(Node{Name: neighbor.Name, Priority: newdist})
			}
		}

		if current == end {
			return dist, parent
		}
	}

	return dist, parent
}

func FindPaths(farm *Farm, start, end string) (map[string]int, Path) {
	dist, parent := Dijkstra(farm, start, end)
	path := Path{}

	if dist[end] == math.MaxInt {
		return nil, nil
	}

	for current := end; current != ""; current = parent[current] {
		path = append(Path{current}, path...)

		if current == start {
			break
		}
	}

	return dist, path
}

func HasDuplicateRoomAcrossPaths(paths [][]string) bool {
	seen := make(map[string]int)

	for i, path := range paths {
		for _, room := range path {
			if prev, exists := seen[room]; exists && prev != i {
				return true
			}
			seen[room] = i
		}
	}
	return false
}
