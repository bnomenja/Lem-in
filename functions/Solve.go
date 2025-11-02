package functions

// import "fmt"

// func EdmondKarp(farm *Farm) ([]Path, []int) {
// 	start := farm.SpecialRooms["start"]
// 	end := farm.SpecialRooms["end"]
// 	pathFound := 0
// 	shortest := []Path{}
// 	foundShortest := false

// 	for {
// 		path := bfs(farm, start, end)
// 		if path == nil {
// 			break
// 		}

// 		if !foundShortest {
// 			shortest = append(shortest, path[1:])
// 			foundShortest = true
// 		}

// 		updateEdges(path, farm)
// 		pathFound++
// 	}

// 	if len(shortest) == 0 {
// 		return shortest, []int{}
// 	}

// 	best := extractPaths(farm, start, end, pathFound)
// 	paths, assigned := findBetterChoice(best, shortest, farm.Antnumber)

// 	return paths, assigned
// }

// func extractPaths(farm *Farm, start, end string, pathNumber int) []Path {
// 	links := buildUsedLinks(farm)
// 	return reconstructPaths(links, start, end, pathNumber)
// }

// func updateEdges(path Path, farm *Farm) {
// 	for i := 0; i < len(path)-1; i++ {
// 		updateOneEdge(path[i], path[i+1], farm)
// 	}
// }

// func bfs(farm *Farm, start, end string) Path {
// 	queue := []string{start}
// 	fmt.Println("queue: ", queue)
// 	visited := map[string]bool{start: true}
// 	parent := map[string]string{}

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		if current == end {
// 			return buildPathFromParents(parent, start, end)
// 		}

// 		exploreNeighbors(farm, current, visited, parent, &queue)
// 	}

// 	return nil
// }
