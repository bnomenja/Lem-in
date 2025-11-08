package functions

import (
	"fmt"
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
			shortest = append(shortest, path[1:])
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

		path := dfs(farm, start, end)
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

		prev := path[i]

		if i > 0 {
			prev = path[i-1]
		}

		current := path[i]
		next := path[i+1]

		currentEdge := farm.Edges[current+"-"+next]
		prevEdge := farm.Edges[prev+"-"+current]
		currentRoom := farm.Rooms[current]

		if current != farm.SpecialRooms["start"] {
			if prevEdge.State == -1 && currentEdge.State == -1 {
				currentRoom.Inpath = false
				farm.Rooms[current] = currentRoom
			} else {
				currentRoom.Inpath = true
				farm.Rooms[current] = currentRoom
			}
		}
	}

	for i := range path {
		if i == len(path)-1 {
			continue
		}

		current := path[i]
		next := path[i+1]
		currentEdge := farm.Edges[current+"-"+next]
		reversedCurrent := farm.Edges[next+"-"+current]

		if currentEdge.State == 1 {
			currentEdge.State = 0
			reversedCurrent.State = -1

			farm.Edges[current+"-"+next] = currentEdge
			farm.Edges[next+"-"+current] = reversedCurrent
		} else {
			currentEdge.State = 1
			reversedCurrent.State = 1

			farm.Edges[current+"-"+next] = currentEdge
			farm.Edges[next+"-"+current] = reversedCurrent
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
		fmt.Println("turn: ", shortTurn)
		return shortest, assignedShort
	}

	fmt.Println("turn: ", turn)

	return best, assigned
}
