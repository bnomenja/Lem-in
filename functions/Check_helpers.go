package functions

import (
	"fmt"
	"strings"
)

func handleSpecialComment(lines []string, idx *int, farm *Farm, foundTunnels *bool, specialType string) error {
	if *idx+1 >= len(lines) {
		return fmt.Errorf("'%s' must not be the last line", specialType)
	}

	nextLine := strings.TrimSpace(lines[*idx+1])

	if msg, ok := isRoomOrTunnel(nextLine, farm, foundTunnels); !ok {
		return fmt.Errorf("after '%s' the next line must be a valid room: %s", specialType, msg)
	}

	name := strings.Fields(nextLine)[0]

	switch specialType {

	case "##start":
		if _, exists := farm.SpecialRooms["start"]; exists {
			return fmt.Errorf("more than one '##start' found")
		}
		farm.SpecialRooms["start"] = name

	case "##end":
		if _, exists := farm.SpecialRooms["end"]; exists {
			return fmt.Errorf("more than one '##end' found")
		}
		farm.SpecialRooms["end"] = name
	}

	*idx++
	return nil
}

func isRoomOrTunnel(line string, farm *Farm,  foundTunnel *bool) (string, bool) {
	if isEmpty(line) {
		return "line is empty", false
	}

	if isComment(line) {
		return "line is a comment", false
	}

	if isRoomLine(line) {
		return handleRoomLine(line, farm, foundTunnel)
	}

	if isTunnelLine(line) {
		return handleTunnelLine(line, farm, foundTunnel)
	}

	return fmt.Sprintf("if you want to comment, you need to use '#' at the begining of the line: %s", line), false
}
