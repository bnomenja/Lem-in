package functions

import (
	"fmt"
	"strconv"
	"strings"
)

type Farm struct {
	Antnumber    int
	Rooms        map[string]*Room
	SpecialRooms map[string]string
	Tunnels      map[string]bool
	Edges        map[string]Edge
}

type Edge struct {
	From  string
	To    string
	State int
}

type Path []string

type Room struct {
	Name   string
	Coord  Position
	Links  []*Room
	Inpath bool
}

type Position struct {
	X int
	Y int
}

type Ant struct {
	Id       int
	Path     Path
	Position int
	Finished bool
}

func validateWords(words []string) error {
	for _, word := range words {
		if word == "" {
			return fmt.Errorf("space between name and coord must be only one")
		}
	}
	return nil
}

func validateRoomName(name string) error {
	if name[0] == 'L' {
		return fmt.Errorf("from room cannot start with 'L'")
	}
	if strings.Contains(name, "-") {
		return fmt.Errorf("room name must not contain from '-")
	}

	return nil
}

func parseCoordinates(xStr, yStr, line string) (int, int, error) {
	x, err1 := strconv.Atoi(xStr)
	y, err2 := strconv.Atoi(yStr)

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("room coords must be integers: %v", line)
	}

	if x < 0 || y < 0 {
		return 0, 0, fmt.Errorf("room coords must be positive: %v", line)
	}

	return x, y, nil
}

func addRoomToFarm(farm *Farm, name string, x, y int) error {
	if _, exists := farm.Rooms[name]; exists {
		return fmt.Errorf("duplicate room: %v", name)
	}

	for _, savedRoom := range farm.Rooms {
		if savedRoom.Coord.X == x && savedRoom.Coord.Y == y {
			return fmt.Errorf("two rooms cannot have the same coord: '%s %v %v' and '%s %v %v'",
				savedRoom.Name, savedRoom.Coord.X, savedRoom.Coord.Y, name, x, y)
		}
	}

	room := &Room{
		Name:  name,
		Coord: Position{X: x, Y: y},
		Links: []*Room{},
	}

	farm.Rooms[name] = room
	return nil
}

func handleTunnelLine(line string, farm *Farm, foundTunnel *bool) (string, bool) {
	if strings.Contains(line, " ") {
		return "tunnel format must not contain spaces", false
	}

	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return "a tunnel links exactly two rooms", false
	}

	from, to := parts[0], parts[1]

	if err := validateTunnelRooms(from, to); err != nil {
		return err.Error(), false
	}

	if err := addTunnelToFarm(farm, from, to, line); err != nil {
		return err.Error(), false
	}

	*foundTunnel = true
	return "", true
}

func validateTunnelRooms(from, to string) error {
	if from == "" || to == "" {
		return fmt.Errorf("tunnel must link two valid room names")
	}
	if from == to {
		return fmt.Errorf("tunnel cannot link from room to itself")
	}
	return nil
}

func addTunnelToFarm(farm *Farm, from, to, line string) error {
	rA, okA := farm.Rooms[from]
	rB, okB := farm.Rooms[to]

	if !okA || !okB {
		return fmt.Errorf("tunnel links non-existing room(s): %v", line)
	}

	k1 := from + "-" + to
	k2 := to + "-" + from

	if farm.Tunnels[k1] || farm.Tunnels[k2] {
		return fmt.Errorf("duplicate tunnel: %v", line)
	}

	farm.Tunnels[k1] = true
	rA.Links = append(rA.Links, rB)
	rB.Links = append(rB.Links, rA)

	farm.Edges[k1] = Edge{
		From:  from,
		To:    to,
		State: 1,
	}

	farm.Edges[k2] = Edge{
		From:  to,
		To:    from,
		State: 1,
	}

	return nil
}

func isRoomLine(line string) bool {
	words := strings.Split(line, " ")
	return len(words) == 3
}

func isTunnelLine(line string) bool {
	return strings.Contains(line, "-")
}

func isComment(line string) bool {
	return len(line) > 0 && line[0] == '#'
}

func isEmpty(line string) bool {
	return line == ""
}

func handleRoomLine(line string, farm *Farm, foundTunnel *bool) (string, bool) {
	words := strings.Split(line, " ")

	if err := validateWords(words); err != nil {
		return err.Error(), false
	}

	if *foundTunnel {
		return "valid format is:\nnumber_of_ants\nthe_rooms\nthe_links", false
	}

	name := words[0]

	if err := validateRoomName(name); err != nil {
		return err.Error(), false
	}

	x, y, err := parseCoordinates(words[1], words[2], line)
	if err != nil {
		return err.Error(), false
	}

	if err := addRoomToFarm(farm, name, x, y); err != nil {
		return err.Error(), false
	}

	return "", true
}
