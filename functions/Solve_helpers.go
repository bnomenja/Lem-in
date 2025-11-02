package functions

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// func findBetterChoice(best, shortest []Path, antNumber int) ([]Path, []int) {
// 	assignedShort, shortTurn := calculateTurns(shortest, antNumber)
// 	assigned, turn := calculateTurns(best, antNumber)

// 	if shortTurn <= turn {
// 		fmt.Println(shortTurn)
// 		return shortest, assignedShort
// 	}

// 	fmt.Println(turn)

// 	return best, assigned
// }

// func calculateTurns(paths []Path, antNumber int) ([]int, int) {
// 	sort.Slice(paths, func(i, j int) bool {
// 		return len(paths[i]) < len(paths[j])
// 	})

// 	assigned := assignAnts(paths, antNumber)

// 	maxTurn := 0
// 	for i := range paths {
// 		turn := len(paths[i]) - 1 + assigned[i]
// 		if turn > maxTurn {
// 			maxTurn = turn
// 		}
// 	}

// 	return assigned, maxTurn
// }

// func buildUsedLinks(farm *Farm) map[string][]string {
// 	links := make(map[string][]string)

// 	for name, edge := range farm.Edges {
// 		if edge.Capacity == 0 {
// 			parts := strings.Split(name, "-")
// 			from, to := parts[0], parts[1]
// 			links[from] = append(links[from], to)
// 		}
// 	}

// 	return links
// }

// func reconstructPaths(links map[string][]string, start, end string, pathNumber int) []Path {
// 	paths := []Path{}
// 	used := make(map[string]bool)

// 	for i := 0; i < pathNumber; i++ {
// 		path := Path{start}
// 		current := start

// 		for current != end {
// 			found := false
// 			for _, to := range links[current] {
// 				tunnel := current + "-" + to

// 				if !used[tunnel] {
// 					used[tunnel] = true
// 					path = append(path, to)
// 					current = to
// 					found = true
// 					break
// 				}
// 			}

// 			if !found {
// 				break
// 			}
// 		}

// 		path = path[1:]
// 		paths = append(paths, path)
// 	}

// 	return paths
// }

// func assignAnts(paths []Path, antNumber int) []int {
// 	pathLen := make([]int, len(paths))

// 	for i, path := range paths {
// 		pathLen[i] = len(path)
// 	}

// 	assigned := make([]int, len(paths))
// 	antsLeft := antNumber

// 	for antsLeft > 0 {
// 		target := findMinLoadPath(pathLen, assigned)
// 		assigned[target]++
// 		antsLeft--
// 	}

// 	return assigned
// }

// func findMinLoadPath(pathLen, assigned []int) int {
// 	target := 0
// 	lowest := pathLen[0] + assigned[0]

// 	for i := 1; i < len(pathLen); i++ {
// 		load := pathLen[i] + assigned[i]
// 		if load <= lowest {
// 			target = i
// 			lowest = load
// 		}
// 	}

// 	return target
// }

// func updateOneEdge(from, to string, farm *Farm) {
// 	forwardKey := from + "-" + to
// 	reverseKey := to + "-" + from

// 	forward := farm.Edges[forwardKey]
// 	forward.Capacity--
// 	farm.Edges[forwardKey] = forward

// 	reverse := farm.Edges[reverseKey]
// 	reverse.Capacity++
// 	farm.Edges[reverseKey] = reverse
// }

// func buildPathFromParents(parent map[string]string, start, end string) Path {
// 	path := Path{}
// 	for room := end; room != ""; room = parent[room] {
// 		path = append([]string{room}, path...)
// 		if room == start {
// 			break
// 		}
// 	}
// 	return path
// }

// func exploreNeighbors(farm *Farm, current string, visited map[string]bool, parent map[string]string, queue *[]string) {
// 	for _, neighbor := range farm.Rooms[current].Links {
// 		if visited[neighbor.Name] {
// 			continue
// 		}

// 		edgeKey := current + "-" + neighbor.Name
// 		if farm.Edges[edgeKey].Capacity > 0 {
// 			visited[neighbor.Name] = true
// 			parent[neighbor.Name] = current
// 			*queue = append(*queue, neighbor.Name)
// 		}
// 	}
// }
