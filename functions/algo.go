package functions

import (
	"math"
)

func bfs(farm *Farm, start, end string) []string {
	parent := map[string]string{}
	queue := []string{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return buildPathfromParent(parent, start, end)
		}

		for _, neighbor := range farm.Rooms[current].Links {
			edge := farm.Edges[current+"-"+neighbor.Name]

			if edge.State != 0 {
				continue
			}

			parent[neighbor.Name] = current
			queue = append(queue, neighbor.Name)
		}
	}

	return nil
}

func dfs(farm *Farm, start, end string) []string {
	current := start
	path := []string{current}
	visited := make(map[string]bool)

	for current != end {
		visited[current] = true
		foundNext := false

		for _, neighbor := range farm.Rooms[current].Links {
			if visited[neighbor.Name] {
				continue
			}

			edge := farm.Edges[current+"-"+neighbor.Name]

			if edge.State != 0 {
				continue
			}

			path = append(path, neighbor.Name)
			current = neighbor.Name
			foundNext = true
			break
		}

		if !foundNext {
			return nil
		}
	}

	if path[len(path)-1] != end {
		return nil
	}

	return path
}

func Dijkstra(farm *Farm, start, end string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	parent := make(map[string]string)
	visited := map[string]bool{}

	for name := range farm.Rooms {
		dist[name] = math.MaxInt
	}
	dist[start] = 0

	queue := queue{}
	queue.Add(Node{Name: start, Priority: 0, OnlyReverse: false})

	for len(queue) > 0 {
		node := queue.Poll()
		current, Value := node.Name, node.Priority

		if dist[current] < Value {
			continue
		}

		visited[current] = true

		for _, neighbor := range farm.Rooms[current].Links {

			key := current + "-" + neighbor.Name
			edge := farm.Edges[key]

			if edge.State == 0 ||
				(node.OnlyReverse && edge.State != -1) ||
				visited[neighbor.Name] {
				continue
			}

			newdist := dist[current] + edge.State
			if newdist < dist[neighbor.Name] {
				parent[neighbor.Name] = current
				dist[neighbor.Name] = newdist

				if edge.State != -1 && neighbor.Inpath {
					queue.Add(Node{Name: neighbor.Name, Priority: newdist, OnlyReverse: true})
					continue

				}

				queue.Add(Node{Name: neighbor.Name, Priority: newdist, OnlyReverse: false})
			}

		}

		if current == end {
			return dist, parent
		}

	}

	return dist, parent
}
