package functions

import "math"

func bfs(farm *Farm) Path {
	start := farm.SpecialRooms["start"]
	end := farm.SpecialRooms["end"]
	parent := map[string]string{}
	visited := map[string]bool{start: true}
	queue := []string{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return buildPathfromParent(parent, start, end)
		}

		for _, neighbor := range farm.Rooms[current].Links {
			key := current + "-" + neighbor.Name
			if visited[neighbor.Name] || farm.Edges[key].Blocked || neighbor.Used {
				continue
			}
			parent[neighbor.Name] = current
			visited[neighbor.Name] = true
			queue = append(queue, neighbor.Name)
		}
	}
	return nil
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

	}
	return dist, parent
}
