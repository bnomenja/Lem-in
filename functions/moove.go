package functions

import (
	"fmt"
	"strings"
)

// MooveAnts simulates the movement of ants through all paths and prints each turn.
func MooveAnts(paths []Path, antNumber int, data string, assigned []int) {
	fmt.Println(strings.TrimSpace(data))
	fmt.Println("")

	ants := make([]Ant, 0)
	finished := 0
	ID := 1

	for finished < antNumber {

		for i := range ants {
			ant := &ants[i]
			last := len(ant.Path) - 1

			if ant.Position < last {
				ant.Position++
				fmt.Printf("L%d-%s ", ant.Id, ant.Path[ant.Position])

			}
			if ant.Position == last && !ant.Finished {
				ant.Finished = true
				finished++
			}
		}

		for i, path := range paths {
			if assigned[i] > 0 {
				newAnt := Ant{
					Id:       ID,
					Path:     path,
					Position: 0,
				}
				ants = append(ants, newAnt)
				fmt.Printf("L%d-%s ", newAnt.Id, newAnt.Path[0])
				ID++
				assigned[i]--
			}
		}
		fmt.Println("")
	}
}
