package functions

import (
	"math"
)

type Node struct {
	Name        string
	Priority    int
	OnlyReverse bool
}

type queue []Node

func Suurballe(farm *Farm) ([]Path, []int) {
	start := farm.SpecialRooms["start"]
	end := farm.SpecialRooms["end"]
	foundShourtest := false
	shortest := []Path{}

	for {
		path := FindPaths(farm, start, end)
		if path == nil {
			break
		}

		if !foundShourtest {
			shortest = append(shortest, path)
			foundShourtest = true
		}

		UpdateGraph(farm, path)
	}

	if !foundShourtest {
		return nil, nil
	}

	paths := MergePaths(farm, start, end)

	best, assigned := findBetterChoice(paths, shortest, farm.Antnumber)

	return best, assigned
}

func MergePaths(farm *Farm, start, end string) []Path {
	merged := []Path{}
	for {

		path := bfs(farm, start, end)
		if path == nil {
			break
		}

		UpdateGraph(farm, path)

		merged = append(merged, path[1:])
	}
	return merged
}

func UpdateGraph(farm *Farm, path Path) {
	for i := range path {
		if i == len(path)-1 {
			continue
		}
		from, to := path[i], path[i+1]

		if from != farm.SpecialRooms["start"] {
			room := farm.Rooms[from]
			room.Inpath = true
			farm.Rooms[from] = room
		}

		edge := farm.Edges[from+"-"+to]
		reverseEdge := farm.Edges[to+"-"+from]

		if edge.State == 1 {
			edge.State = 0
			reverseEdge.State = -1

			farm.Edges[from+"-"+to] = edge
			farm.Edges[to+"-"+from] = reverseEdge
		} else {
			edge.State = 1
			reverseEdge.State = 1

			farm.Edges[from+"-"+to] = edge
			farm.Edges[to+"-"+from] = reverseEdge
		}
	}
}

func FindPaths(farm *Farm, start, end string) Path {
	dist, parent := Dijkstra(farm, start, end)
	if dist[end] == math.MaxInt {
		return nil
	}

	path := buildPathfromParent(parent, start, end)
	return path
}

func findBetterChoice(best, shortest []Path, antNumber int) ([]Path, []int) {
	assignedShort, shortTurn := CalculateTurns(shortest, antNumber)
	assigned, turn := CalculateTurns(best, antNumber)

	if shortTurn <= turn {
		return shortest, assignedShort
	}

	return best, assigned
}
