package functions

import (
	"fmt"
	"math"
)

type Node struct {
	Name     string
	Priority int
}

type queue []Node

func Suurballe(farm *Farm) ([]Path, []int) {
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
		return nil, nil
	}

	shortest := paths[0][1:]
	paths = MergePaths(paths, farm)
	fmt.Println(paths)

	best, assigned := findBetterChoice(paths, []Path{shortest}, farm.Antnumber)

	return best, assigned
}

func MergePaths(paths []Path, farm *Farm) []Path {
	merged := []Path{}
	BlockUselessEdges(farm, paths)

	for {
		path := bfs(farm)
		if path == nil {
			break
		}
		if len(path) == 2 {
			return []Path{path[1:]}
		}
		for _, name := range path[1 : len(path)-1] {
			room := farm.Rooms[name]
			room.Used = true
			farm.Rooms[name] = room
		}
		merged = append(merged, path[1:])
	}
	return merged
}

func UpdateGraph(farm *Farm, shortest Path, dist map[string]int) {
	for key, edge := range farm.Edges {
		from, to := edge.From, edge.To
		newWeight := edge.Weight + dist[from] - dist[to]

		farm.Edges[key] = Edge{
			From:    from,
			To:      to,
			Weight:  newWeight,
			Blocked: edge.Blocked,
		}
	}

	for i := 0; i < len(shortest)-1; i++ {
		from, to := shortest[i], shortest[i+1]
		edge := farm.Edges[to+"-"+from]
		edge.Weight = 0
		farm.Edges[to+"-"+from] = edge
		delete(farm.Edges, from+"-"+to)
	}
}

func FindPaths(farm *Farm, start, end string) (map[string]int, Path) {
	dist, parent := Dijkstra(farm, start, end)
	if dist[end] == math.MaxInt {
		return nil, nil
	}
	path := buildPathfromParent(parent, start, end)
	return dist, path
}

func findBetterChoice(best, shortest []Path, antNumber int) ([]Path, []int) {
	if len(shortest[0]) == 1 {
		return shortest, []int{antNumber}
	}

	assignedShort, shortTurn := calculateTurns(shortest, antNumber)
	assigned, turn := calculateTurns(best, antNumber)

	if shortTurn <= turn {
		return shortest, assignedShort
	}

	return best, assigned
}
